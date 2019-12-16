package endpoints

import (
	"context"
	"encoding/json"

	"github.com/adlio/trello"
	"github.com/open-integration/core/pkg/logger"
)

func UnmarshalArchiveCardsArguments(data []byte) (ArchiveCardsArguments, error) {
	var r ArchiveCardsArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ArchiveCardsArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ArchiveCardsArguments struct {
	App     string   `json:"App"`     // Trello Application ID
	CardIDs []string `json:"CardIDs"` // IDs to archive
	Token   string   `json:"Token"`   // Trello Token
}

type ArchiveCardsReturns struct{}

func UnmarshalArchiveCardsReturns(data []byte) (ArchiveCardsReturns, error) {
	var r ArchiveCardsReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ArchiveCardsReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func ArchiveCards(context context.Context, log logger.Logger, args *ArchiveCardsArguments) (*ArchiveCardsReturns, error) {
	client := trello.NewClient(args.App, args.Token)

	for _, id := range args.CardIDs {
		if id == "" {
			continue
		}
		card, err := client.GetCard(id, trello.Defaults())
		if err != nil {
			log.Error("Failed to get card", "card", id, "error", err.Error())
			return nil, err
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
	return &ArchiveCardsReturns{}, nil
}
