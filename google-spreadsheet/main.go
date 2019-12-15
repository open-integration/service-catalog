package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/grpc"

	"github.com/olegsu/service-catalog/google-spreadsheet/configs/endpoints"
	"github.com/open-integration/core/pkg/logger"

	api "github.com/open-integration/core/pkg/api/v1"
)

type (
	Service struct {
		logger logger.Logger
	}

	RemoteRow struct {
		UpdatedAt   time.Time
		RemoteIndex int
		Row         *Row
	}
)

func main() {
	service := &Service{
		logger: logger.New(nil),
	}
	runServer(context.Background(), service, os.Getenv("PORT"))
}

func (s *Service) Init(context context.Context, req *api.InitRequest) (*api.InitResponse, error) {
	schemas := map[string]string{}
	for k, v := range endpoints.TemplatesMap() {
		schemas[k] = v
	}
	return &api.InitResponse{
		JsonSchemas: schemas,
	}, nil
}

func (s *Service) Call(context context.Context, req *api.CallRequest) (*api.CallResponse, error) {
	res := &api.CallResponse{
		Status:  api.Status_OK,
		Payload: "{}",
	}
	log := logger.New(&logger.Options{
		FilePath: req.Fd,
	})
	args, err := UnmarshalArguments([]byte(req.Arguments))
	if err != nil {
		log.Error("Failed to load convert arguments string", "error", err.Error())
		res.Status = api.Status_Error
		res.Error = err.Error()
		return res, nil
	}
	if req.Endpoint == "Upsert" {
		log.Debug("Args", "service-account", args.ServiceAccount, "spreadsheet-id", args.SpreadsheetID)
		sp, err := connect(args.ServiceAccount, args.SpreadsheetID, log)
		if err != nil {
			log.Error("Failed to connect to google", "error", err.Error())
			res.Status = api.Status_Error
			res.Error = err.Error()
			return res, nil
		}
		log.Debug("Connected to Google")

		r, err := sp.Spreadsheets.Values.Get(args.SpreadsheetID, "A:C").Do()
		if err != nil {
			log.Error("Failed get values from speadsheet", "error", err.Error())
			res.Status = api.Status_Error
			res.Error = err.Error()
			return res, nil
		}

		remoteRows := map[string]*RemoteRow{}

		if len(r.Values) > 0 {
			for index, ra := range r.Values[1:] { // skip header
				t, _ := time.Parse("02-01-2006 15:04:05", ra[2].(string))
				remoteRows[ra[0].(string)] = &RemoteRow{
					UpdatedAt:   t,
					RemoteIndex: index + 2,
				}
			}
		}

		rowsToAdd := [][]interface{}{}
		rowsToUpdate := map[string]*RemoteRow{}

		for _, row := range args.Rows {
			remoteRow, exist := remoteRows[*row.ID]
			t, err := time.Parse("02-01-2006 15:04:05", row.Data[1].(string))
			if err != nil {
				log.Error("Failed parse time from cell", "time", row.Data[1], "error", err.Error())
				continue
			}
			if !exist {
				r := []interface{}{
					row.ID,
				}
				r = append(r, row.Data...)
				rowsToAdd = append(rowsToAdd, r)
			} else if exist && t != remoteRow.UpdatedAt {
				remoteRow.Row = &row
				rowsToUpdate[*row.ID] = remoteRow
			}
		}

		for _, rowToUpdate := range rowsToUpdate {
			data := []interface{}{
				rowToUpdate.Row.ID,
			}
			data = append(data, rowToUpdate.Row.Data...)
			vr := &sheets.ValueRange{
				Values: [][]interface{}{
					data,
				},
			}
			rangeStr := fmt.Sprintf("%d:%d", rowToUpdate.RemoteIndex, rowToUpdate.RemoteIndex)
			log.Debug("Updating row", "range", rangeStr, "data", data)
			_, err := sp.Spreadsheets.Values.Update(args.SpreadsheetID, rangeStr, vr).ValueInputOption("RAW").Do()
			if err != nil {
				log.Error("Failed to update row", "range", rangeStr, "error", err.Error())
			}
		}

		log.Debug("Adding new rows", "len", len(rowsToAdd))
		_, err = sp.Spreadsheets.Values.Append(args.SpreadsheetID, "1:1", &sheets.ValueRange{
			Values: rowsToAdd,
		}).ValueInputOption("RAW").Do()
		if err != nil {
			log.Error("Failed to add new row", "error", err.Error())
		}

		_, err = sp.Spreadsheets.BatchUpdate(args.SpreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				&sheets.Request{
					ClearBasicFilter: &sheets.ClearBasicFilterRequest{
						SheetId: 0,
					},
				},
				&sheets.Request{
					SetBasicFilter: &sheets.SetBasicFilterRequest{
						Filter: &sheets.BasicFilter{
							Range: &sheets.GridRange{
								StartColumnIndex: 0,
								EndColumnIndex:   7,
							},
						},
					},
				},
			},
		}).Do()

		if err != nil {
			log.Error("Failed to run batch update call", "error", err.Error())
		}

	}
	return res, nil
}

func runServer(ctx context.Context, v1API api.ServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("starting gRPC server, listening on port %s\n", port)
	err = server.Serve(listen)
	if err != nil {
		log.Printf("Error starting gRPC server: %s", err.Error())
		os.Exit(1)
	}
	return nil
}

func connect(serviceAccountFilePath string, spreadsheetID string, log logger.Logger) (*sheets.Service, error) {
	b, err := ioutil.ReadFile(serviceAccountFilePath)
	if err != nil {
		return nil, err
	}

	driveScopes := []string{
		drive.DriveScope,
		drive.DriveAppdataScope,
		drive.DriveFileScope,
		drive.DriveMetadataScope,
		drive.DriveMetadataReadonlyScope,
		drive.DrivePhotosReadonlyScope,
		drive.DriveReadonlyScope,
		drive.DriveScriptsScope,
	}
	config, err := google.JWTConfigFromJSON(b, driveScopes...)
	if err != nil {
		return nil, err
	}
	client := config.Client(context.Background())
	return sheets.New(client)
}

func load(j string) ([]*Row, error) {
	cards := []*Row{}
	err := json.Unmarshal([]byte(j), &cards)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
