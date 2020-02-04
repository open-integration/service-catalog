package addcard

import (
	"context"

	"github.com/adlio/trello"
	"github.com/open-integration/core/pkg/logger"
)

type (
	AddcardOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *AddcardArguments
	}
)

func Addcard(opt AddcardOptions) (*AddcardReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	args := opt.Arguments
	client := trello.NewClient(args.App, args.Token)

	card := &trello.Card{
		Name:     opt.Arguments.Name,
		IDList:   opt.Arguments.List,
		IDLabels: opt.Arguments.Labels,
	}
	if opt.Arguments.Description != nil {
		card.Desc = *opt.Arguments.Description
	}

	err := client.CreateCard(card, trello.Defaults())
	if err != nil {
		log.Error("Failed to create card", "boardID", args.Board, "error", err.Error())
		return nil, err
	}
	log.Debug("card added", "boardID", args.Board)
	return &AddcardReturns{}, nil
}
