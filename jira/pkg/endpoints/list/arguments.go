// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    listArguments, err := UnmarshalListArguments(bytes)
//    bytes, err = listArguments.Marshal()

package list

import "encoding/json"

func UnmarshalListArguments(data []byte) (ListArguments, error) {
	var r ListArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ListArguments struct {
	APIToken    string  `json:"API_Token"`            // Jira API token
	Endpoint    string  `json:"Endpoint"`             // Jira endpoint
	Jql         *string `json:"JQL,omitempty"`        // Jira Query Language string
	QueryFields *string `json:"QueryFields,omitempty"`// Jira query fields to include in response
	Sort        *string `json:"Sort,omitempty"`       // Jira sort properties
	User        string  `json:"User"`                 // IDs to archive
}
