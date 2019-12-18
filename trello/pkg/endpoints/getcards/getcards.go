package getcards

import (
	"context"
	"encoding/json"

	"github.com/adlio/trello"
	"github.com/open-integration/core/pkg/logger"
)

// Getcards return all the cards in given board id
func Getcards(context context.Context, log logger.Logger, args *GetcardsArguments) (*GetcardsReturns, error) {
	client := trello.NewClient(args.App, args.Token)
	board, err := client.GetBoard(args.Board, trello.Defaults())
	if err != nil {
		log.Error("Failed to get boards", "boardID", args.Board, "error", err.Error())
		return nil, err
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		log.Error("Failed to get board lists", "boardID", args.Board, "error", err.Error())
		return nil, err
	}
	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		log.Error("Failed to get board cards", "boardID", args.Board, "error", err.Error())
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
	j, err := json.Marshal(cards)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalGetcardsReturns(j)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
