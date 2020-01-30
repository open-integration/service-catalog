package message

import (
	"bytes"
	"context"
	"net/http"
	"strings"

	"github.com/open-integration/core/pkg/logger"
)

type (
	MessageOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *MessageArguments
	}
)

func Message(opt MessageOptions) (*MessageReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	log.Debug("dummy log")

	var buffer bytes.Buffer

	buffer.WriteString(`{ "text": "`)
	buffer.WriteString(opt.Arguments.Message)
	buffer.WriteString(`"}`)

	log.Info("Sending message", "url", opt.Arguments.WebhookURL)

	res, err := http.Post(opt.Arguments.WebhookURL, "application/x-www-form-urlencoded", strings.NewReader(buffer.String()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &MessageReturns{}, nil
}
