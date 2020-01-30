package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coryb/oreo"
	"github.com/go-jira/jira"
	"github.com/open-integration/core/pkg/logger"
)

type (
	ListOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *ListArguments
	}
)

func List(opt ListOptions) (*ListReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	o := oreo.New().WithPreCallback(extendRequestWithAuthenticatedUser(opt.Arguments.User, opt.Arguments.APIToken))
	log.Debug("Sending search request", "endpoint", opt.Arguments.Endpoint)
	searchOpt := &jira.SearchOptions{}
	if opt.Arguments.Jql != nil {
		searchOpt.Query = *opt.Arguments.Jql
	}
	if opt.Arguments.QueryFields != nil {
		searchOpt.QueryFields = *opt.Arguments.QueryFields
	}
	if opt.Arguments.Sort != nil {
		searchOpt.Sort = *opt.Arguments.Sort
	}
	resp, err := jira.Search(o, opt.Arguments.Endpoint, searchOpt)
	if err != nil {
		return nil, fmt.Errorf("Failed to run search query %w", err)
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal response")
	}
	listResp, err := UnmarshalListReturns(b)
	return &listResp, err
}
