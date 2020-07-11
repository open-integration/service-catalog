// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    getEventsArguments, err := UnmarshalGetEventsArguments(bytes)
//    bytes, err = getEventsArguments.Marshal()

package getEvents

import "encoding/json"

func UnmarshalGetEventsArguments(data []byte) (GetEventsArguments, error) {
	var r GetEventsArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetEventsArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetEventsArguments struct {
	CalendarID              string         `json:"CalendarID"`                        // Google calendar-id
	ICalUID                 *string        `json:"ICalUID,omitempty"`                 // Specifies event ID in the iCalendar format to be included in the response. Optional.
	MaxAttendees            *int64         `json:"MaxAttendees,omitempty"`            // The maximum number of attendees to include in the response. If there are more than the; specified number of attendees, only the participant is returned. Optional.
	MaxResults              *int64         `json:"MaxResults,omitempty"`              // Maximum number of events returned on one result page. The number of events in the; resulting page may be less than this value, or none at all, even if there are more events; matching the query. Incomplete pages can be detected by a non-empty nextPageToken field; in the response. By default the value is 250 events. The page size can never be larger; than 2500 events. Optional.
	OrderBy                 *OrderBy       `json:"OrderBy,omitempty"`                 // The order of the events returned in the result. Optional. The default is an unspecified,; stable order.
	PrivateExtendedProperty *string        `json:"PrivateExtendedProperty,omitempty"` // Extended properties constraint specified as propertyName=value. Matches only private; properties. This parameter might be repeated multiple times to return events that match; all given constraints.
	Q                       *string        `json:"Q,omitempty"`                       // Free text search terms to find events that match these terms in any field, except for; extended properties. Optional.
	ServiceAccount          ServiceAccount `json:"ServiceAccount"`                    // service-account
	SharedExtendedProperty  *string        `json:"SharedExtendedProperty,omitempty"`  // Extended properties constraint specified as propertyName=value. Matches only shared; properties. This parameter might be repeated multiple times to return events that match; all given constraints.
	ShowDeleted             *bool          `json:"ShowDeleted,omitempty"`             // Whether to include deleted events (with status equals "cancelled") in the result.; Cancelled instances of recurring events (but not the underlying recurring event) will; still be included if showDeleted and singleEvents are both False. If showDeleted and; singleEvents are both True, only single instances of deleted events (but not the; underlying recurring events) are returned. Optional. The default is False.
	ShowHiddenInvitations   *bool          `json:"ShowHiddenInvitations,omitempty"`   // Whether to include hidden invitations in the result. Optional. The default is False.
	SingleEvents            *bool          `json:"SingleEvents,omitempty"`            // Whether to expand recurring events into instances and only return single one-off events; and instances of recurring events, but not the underlying recurring events themselves.; Optional. The default is False.
	TimeMax                 *string        `json:"TimeMax,omitempty"`                 // Upper bound (exclusive) for an event's start time to filter by. Optional. The default is; not to filter by start time. Must be an RFC3339 timestamp with mandatory time zone; offset, for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z. Milliseconds may be; provided but are ignored. If timeMin is set, timeMax must be greater than timeMin.
	TimeMin                 *string        `json:"TimeMin,omitempty"`                 // Lower bound (exclusive) for an event's end time to filter by. Optional. The default is; not to filter by end time. Must be an RFC3339 timestamp with mandatory time zone offset,; for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z. Milliseconds may be; provided but are ignored. If timeMax is set, timeMin must be smaller than timeMax.
	TimeZone                *string        `json:"TimeZone,omitempty"`                // Time zone used in the response. Optional. The default is the time zone of the calendar.
	UpdatedMin              *string        `json:"UpdatedMin,omitempty"`              // Lower bound for an event's last modification time (as a RFC3339 timestamp) to filter by.; When specified, entries deleted since this time will always be included regardless of; showDeleted. Optional. The default is not to filter by last modification time.
}

// service-account
type ServiceAccount struct {
	AuthProviderX509CERTURL string `json:"auth_provider_x509_cert_url"`
	AuthURI                 string `json:"auth_uri"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	ClientX509CERTURL       string `json:"client_x509_cert_url"`
	PrivateKey              string `json:"private_key"`
	PrivateKeyID            string `json:"private_key_id"`
	ProjectID               string `json:"project_id"`
	TokenURI                string `json:"token_uri"`
	Type                    string `json:"type"`
}

// The order of the events returned in the result. Optional. The default is an unspecified,
// stable order.
type OrderBy string

const (
	StartTime OrderBy = "startTime"
	Updated   OrderBy = "updated"
)
