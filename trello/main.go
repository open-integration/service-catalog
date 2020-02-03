package main

import (
	"context"
	"fmt"

	"net"
	"os"
	"os/signal"

	"github.com/gobuffalo/packr"
	"google.golang.org/grpc"

	"github.com/open-integration/core/pkg/logger"

	api "github.com/open-integration/core/pkg/api/v1"

	"github.com/open-integration/service-catalog/trello/pkg/endpoints/addcard"
	"github.com/open-integration/service-catalog/trello/pkg/endpoints/archivecards"

	"github.com/open-integration/service-catalog/trello/pkg/endpoints/getcards"
)

type (
	Service struct {
		logger logger.Logger
		box    packr.Box
	}
)

func main() {
	service := &Service{
		logger: logger.New(&logger.Options{
			LogToStdOut: true,
		}),
		box: packr.NewBox("./configs"),
	}
	runServer(context.Background(), service, os.Getenv("PORT"), service.logger)
}

func (s *Service) Init(context context.Context, req *api.InitRequest) (*api.InitResponse, error) {
	schemas := map[string]string{}
	for _, v := range s.box.List() {
		schema, _ := s.box.FindString(v)
		schemas[v] = schema
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

	response := &api.CallResponse{}

	switch req.Endpoint {

	case "archivecards":
		args, err := archivecards.UnmarshalArchivecardsArguments([]byte(req.Arguments))
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		res, err := archivecards.Archivecards(context, log, &args)
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		payload, err := res.Marshal()
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		response.Status = api.Status_OK
		response.Payload = string(payload)
		return response, nil

	case "getcards":
		args, err := getcards.UnmarshalGetcardsArguments([]byte(req.Arguments))
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		res, err := getcards.Getcards(context, log, &args)
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		payload, err := res.Marshal()
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		response.Status = api.Status_OK
		response.Payload = string(payload)
		return response, nil
	case "addcard":
		args, err := addcard.UnmarshalAddcardArguments([]byte(req.Arguments))
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		res, err := addcard.Addcard(addcard.AddcardOptions{
			Context:   context,
			LoggerFD:  req.Fd,
			Arguments: &args,
		})
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		payload, err := res.Marshal()
		if resp := buildErrorResponse(err); resp != nil {
			return resp, nil
		}
		response.Status = api.Status_OK
		response.Payload = string(payload)
		return response, nil

	}
	return buildErrorResponse(fmt.Errorf("Endpoint %s not found", req.Endpoint)), nil
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

func buildErrorResponse(err error) *api.CallResponse {
	if err != nil {
		return &api.CallResponse{
			Error:  err.Error(),
			Status: api.Status_Error,
		}
	}
	return nil
}
