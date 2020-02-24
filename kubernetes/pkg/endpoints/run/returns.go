// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    runReturns, err := UnmarshalRunReturns(bytes)
//    bytes, err = runReturns.Marshal()

package run

import "encoding/json"

type RunReturns map[string]interface{}

func UnmarshalRunReturns(data []byte) (RunReturns, error) {
	var r RunReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RunReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
