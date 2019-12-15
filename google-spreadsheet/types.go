package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//

import "encoding/json"

func UnmarshalArguments(data []byte) (Arguments, error) {
	var r Arguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Arguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Arguments struct {
	Rows           []Row  `json:"Rows"`           // Do upsert
	ServiceAccount string `json:"ServiceAccount"` // Path to service-account.json file
	SpreadsheetID  string `json:"SpreadsheetID"`  // ID of the spreadsheet
}

type Row struct {
	Data []interface{} `json:"Data"`
	ID   *string       `json:"ID,omitempty"`
}
