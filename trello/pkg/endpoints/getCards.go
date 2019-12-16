package endpoints

import (
	"context"
	"encoding/json"

	"github.com/adlio/trello"
	"github.com/open-integration/core/pkg/logger"
)

func UnmarshalGetCardsArguments(data []byte) (GetCardsArguments, error) {
	var r GetCardsArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetCardsArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetCardsArguments struct {
	App   string `json:"App"`   // Trello Application ID
	Board string `json:"Board"` // Trello Board ID
	Token string `json:"Token"` // Trello Token
}

type GetCardsReturns []Card

func UnmarshalGetCardsReturns(data []byte) (GetCardsReturns, error) {
	var r GetCardsReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetCardsReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Card struct {
	Badges                Badges        `json:"badges"`
	Board                 interface{}   `json:"Board"`
	Closed                bool          `json:"closed"`
	DateLastActivity      string        `json:"dateLastActivity"`
	Desc                  string        `json:"desc"`
	Due                   *string       `json:"due"`
	DueComplete           bool          `json:"dueComplete"`
	Email                 string        `json:"email"`
	ID                    string        `json:"id"`
	IDAttachmentCover     string        `json:"idAttachmentCover"`
	IDBoard               string        `json:"idBoard"`
	IDCheckLists          []interface{} `json:"idCheckLists"`
	IDLabels              []string      `json:"idLabels"`
	IDList                string        `json:"idList"`
	IDShort               float64       `json:"idShort"`
	Labels                []Label       `json:"labels"`
	List                  List          `json:"List"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Name                  string        `json:"name"`
	Pos                   float64       `json:"pos"`
	ShortLink             string        `json:"shortLink"`
	ShortURL              string        `json:"shortUrl"`
	Subscribed            bool          `json:"subscribed"`
	URL                   string        `json:"url"`
}

type Badges struct {
	Attachments        float64 `json:"attachments"`
	CheckItems         float64 `json:"checkItems"`
	CheckItemsChecked  float64 `json:"checkItemsChecked"`
	Comments           float64 `json:"comments"`
	Description        bool    `json:"description"`
	Subscribed         bool    `json:"subscribed"`
	ViewingMemberVoted bool    `json:"viewingMemberVoted"`
	Votes              float64 `json:"votes"`
}

type Label struct {
	Color   string  `json:"color"`
	ID      string  `json:"id"`
	IDBoard string  `json:"idBoard"`
	Name    string  `json:"name"`
	Uses    float64 `json:"uses"`
}

type List struct {
	Closed     bool    `json:"closed"`
	ID         string  `json:"id"`
	IDBoard    string  `json:"idBoard"`
	Name       string  `json:"name"`
	Pos        float64 `json:"pos"`
	Subscribed bool    `json:"subscribed"`
}

// GetCards return all the cards in given board id
func GetCards(context context.Context, log logger.Logger, args *GetCardsArguments) (*GetCardsReturns, error) {
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
	res, err := UnmarshalGetCardsReturns(j)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
