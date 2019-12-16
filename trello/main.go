package main

import (
	"context"
	"fmt"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/open-integration/core/pkg/logger"

	ep "github.com/open-integration/service-catalog/trello/configs/endpoints"

	api "github.com/open-integration/core/pkg/api/v1"

	"github.com/open-integration/service-catalog/trello/pkg/endpoints"
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
	for k, v := range ep.TemplatesMap() {
		schemas[k] = v
	}
	return &api.InitResponse{
		JsonSchemas: schemas,
	}, nil
}

func (s *Service) Call(context context.Context, req *api.CallRequest) (*api.CallResponse, error) {
	s.logger.Debug("Request", "endpoint", req.Endpoint)
	log := logger.New(&logger.Options{
		FilePath: req.Fd,
	})

	res := &api.CallResponse{}

	switch req.Endpoint {
	case "GetCards":
		args, err := endpoints.UnmarshalGetCardsArguments([]byte(req.Arguments))
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		cards, err := endpoints.GetCards(context, log, &args)
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		payload, err := cards.Marshal()
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		res.Status = api.Status_OK
		res.Payload = string(payload)
		return res, nil

	case "ArchiveCard":
		args, err := endpoints.UnmarshalArchiveCardsArguments([]byte(req.Arguments))
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		cards, err := endpoints.ArchiveCards(context, log, &args)
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		payload, err := cards.Marshal()
		if resp := buildErrorResponse(err, log); resp != nil {
			return resp, nil
		}

		res.Status = api.Status_OK
		res.Payload = string(payload)
		return res, nil
	}
	return buildErrorResponse(fmt.Errorf("Endpoint %s not found", req.Endpoint), log), nil
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

func buildErrorResponse(err error, log logger.Logger) *api.CallResponse {
	if err != nil {
		log.Error(err.Error())
		return &api.CallResponse{
			Error:  err.Error(),
			Status: api.Status_Error,
		}
	}
	return nil
}
