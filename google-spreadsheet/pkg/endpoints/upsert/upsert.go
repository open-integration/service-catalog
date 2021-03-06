package upsert

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/open-integration/core/pkg/logger"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
)

type (
	RemoteRow struct {
		UpdatedAt   time.Time
		RemoteIndex int
		Row         *Row
	}
)

func Upsert(context context.Context, log logger.Logger, args *UpsertArguments) (*UpsertReturns, error) {
	sp, err := connect(args.ServiceAccount, args.SpreadsheetID, log)
	if err != nil {
		return nil, err
	}
	log.Debug("Connected to Google")
	r, err := sp.Spreadsheets.Values.Get(args.SpreadsheetID, "A:C").Do()
	if err != nil {
		return nil, err
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
			*rowToUpdate.Row.ID,
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
		return nil, err
	}
	return &UpsertReturns{}, nil
}

func connect(serviceAccount ServiceAccount, spreadsheetID string, log logger.Logger) (*sheets.Service, error) {
	b, err := json.Marshal(serviceAccount)

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
