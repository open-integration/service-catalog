// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    upsertArguments, err := UnmarshalUpsertArguments(bytes)
//    bytes, err = upsertArguments.Marshal()

package upsert

import "encoding/json"

func UnmarshalUpsertArguments(data []byte) (UpsertArguments, error) {
	var r UpsertArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UpsertArguments struct {
	Rows           []Row  `json:"Rows"`          // Do upsert
	ServiceAccount string `json:"ServiceAccount"`// Path to service-account.json file
	SpreadsheetID  string `json:"SpreadsheetID"` // ID of the spreadsheet
}

type Row struct {
	Data []interface{} `json:"Data"`        
	ID   *string       `json:"ID,omitempty"`
}
