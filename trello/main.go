package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/adlio/trello"
	"github.com/olegsu/openc/pkg/logger"

	"github.com/olegsu/openc-services/trello/configs/endpoints"

	api "github.com/olegsu/openc/pkg/api/v1"
)

type (
	Service struct {
		logger logger.Logger
	}
)

func main() {
	service := &Service{
		logger: logger.New(nil),
	}
	runServer(context.Background(), service, os.Getenv("PORT"), service.logger)
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
	switch req.Endpoint {
	case "GetCards":
		return s.getCardsEndpoint(context, req), nil
	case "ArchiveCard":
		return s.archiveCardEndpoint(context, req), nil
	}

	return &api.CallResponse{
		Error:  fmt.Sprintf("Endpoint %s not found", req.Endpoint),
		Status: api.Status_OK,
	}, nil
}

func runServer(ctx context.Context, v1API api.ServiceServer, port string, log logger.Logger) error {
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
			log.Debug("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Debug("starting gRPC server", "port", port)
	err = server.Serve(listen)
	if err != nil {
		log.Debug("Error starting gRPC server", "error", err.Error())
		os.Exit(1)
	}
	return nil
}

func (s *Service) getCardsEndpoint(context context.Context, req *api.CallRequest) *api.CallResponse {
	log := logger.New(&logger.Options{
		FilePath: req.Fd,
	})
	res := &api.CallResponse{}
	args := req.Arguments
	cards, err := s.request(args["App"], args["Token"], args["Board"], log)
	if err != nil {
		res.Status = api.Status_Error
		res.Error = err.Error()
	} else {
		data, err := json.Marshal(cards)
		if err != nil {
			res.Status = api.Status_Error
			res.Error = err.Error()
		} else {
			res.Status = api.Status_OK
			res.Payload = string(data)
		}
	}
	return res
}

func (s *Service) archiveCardEndpoint(context context.Context, req *api.CallRequest) *api.CallResponse {
	res := &api.CallResponse{
		Status: api.Status_OK,
	}
	log := logger.New(&logger.Options{
		FilePath: req.Fd,
	})

	args := req.Arguments
	client := trello.NewClient(args["App"], args["Token"])

	cardsIds := strings.Split(req.Arguments["CardIDs"], ",")
	for _, id := range cardsIds {
		if id == "" {
			continue
		}
		card, err := client.GetCard(id, trello.Defaults())
		if err != nil {
			log.Error("Failed to get card", "card", id, "error", err.Error())
			res.Status = api.Status_Error
			res.Error = err.Error()
			return res
		}
		err = card.Update(trello.Arguments{
			"closed": "true",
		})
		if err != nil {
			log.Error("Failed to archive card", "card", id, "error", err.Error())
			continue
		}
		log.Debug("Card archived", "Card", card.ID)
	}
	return res
}

func (s *Service) request(app string, token string, boardID string, log logger.Logger) ([]*trello.Card, error) {
	client := trello.NewClient(app, token)
	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		log.Error("Failed to get boards", "boardID", boardID, "error", err.Error())
		return nil, err
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		log.Error("Failed to get board lists", "boardID", boardID, "error", err.Error())
		return nil, err
	}
	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		log.Error("Failed to get board cards", "boardID", boardID, "error", err.Error())
		return nil, err
	}

	for _, card := range cards {
		var list *trello.List
		for _, l := range lists {
			if card.IDList == l.ID {
				list = l
			}
		}
		card.List = list
		log.Debug("Card", "card", card.Name)
	}
	return cards, nil
}
