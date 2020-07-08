package getEvents

import (
	"context"
	"encoding/json"

	"github.com/open-integration/core/pkg/logger"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type (
	GetEventsOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *GetEventsArguments
	}
)

func GetEvents(opt GetEventsOptions) (*GetEventsReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	svc, err := connect(opt.Arguments.ServiceAccount, log)
	if err != nil {
		return nil, err
	}
	log.Debug("Connected to Google Calendar")
	req := svc.Events.List(opt.Arguments.CalendarID)

	if opt.Arguments.SingleEvents != nil {
		req.SingleEvents(*opt.Arguments.SingleEvents)
	}
	if opt.Arguments.ICalUID != nil {
		req.ICalUID(*opt.Arguments.ICalUID)
	}

	if opt.Arguments.MaxAttendees != nil {
		req.MaxAttendees(*opt.Arguments.MaxAttendees)
	}

	if opt.Arguments.MaxResults != nil {
		req.MaxResults(*opt.Arguments.MaxResults)
	}

	if opt.Arguments.OrderBy != nil {
		req.OrderBy(string(*opt.Arguments.OrderBy))
	}

	if opt.Arguments.PrivateExtendedProperty != nil {
		req.PrivateExtendedProperty(*opt.Arguments.PrivateExtendedProperty)
	}

	if opt.Arguments.Q != nil {
		req.Q(*opt.Arguments.Q)
	}

	if opt.Arguments.SharedExtendedProperty != nil {
		req.SharedExtendedProperty(*opt.Arguments.SharedExtendedProperty)
	}

	if opt.Arguments.ShowDeleted != nil {
		req.ShowDeleted(*opt.Arguments.ShowDeleted)
	}

	if opt.Arguments.ShowHiddenInvitations != nil {
		req.ShowHiddenInvitations(*opt.Arguments.ShowHiddenInvitations)
	}

	if opt.Arguments.SingleEvents != nil {
		req.SingleEvents(*opt.Arguments.SingleEvents)
	}

	if opt.Arguments.TimeMax != nil {
		req.TimeMax(*opt.Arguments.TimeMax)
	}

	if opt.Arguments.TimeMin != nil {
		req.TimeMin(*opt.Arguments.TimeMin)
	}

	if opt.Arguments.TimeZone != nil {
		req.TimeZone(*opt.Arguments.TimeZone)
	}

	res, err := req.Context(opt.Context).Do()
	if err != nil {
		return nil, err
	}
	result := []Event{}
	for _, ev := range res.Items {
		d, err := ev.MarshalJSON()
		if err != nil {
			return nil, err
		}
		event := Event{}
		if err := json.Unmarshal(d, &event); err != nil {
			return nil, err
		}
		result = append(result, event)
	}

	return &GetEventsReturns{
		Events: result,
	}, nil
}

func connect(serviceAccount ServiceAccount, log logger.Logger) (*calendar.Service, error) {
	b, err := json.Marshal(serviceAccount)

	scopes := []string{
		calendar.CalendarScope,
		calendar.CalendarEventsReadonlyScope,
	}

	config, err := google.JWTConfigFromJSON(b, scopes...)
	if err != nil {
		return nil, err
	}
	client := config.Client(context.Background())
	return calendar.New(client)
}
