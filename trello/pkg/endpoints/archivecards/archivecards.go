package archivecards

import (
	"context"

	"github.com/adlio/trello"
	"github.com/open-integration/core/pkg/logger"
)

func Archivecards(context context.Context, log logger.Logger, args *ArchivecardsArguments) (*ArchivecardsReturns, error) {
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
	return &ArchivecardsReturns{}, nil
}
