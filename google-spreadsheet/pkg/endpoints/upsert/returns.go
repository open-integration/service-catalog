// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    upsertReturns, err := UnmarshalUpsertReturns(bytes)
//    bytes, err = upsertReturns.Marshal()

package upsert

import "encoding/json"

type UpsertReturns map[string]interface{}

func UnmarshalUpsertReturns(data []byte) (UpsertReturns, error) {
	var r UpsertReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
