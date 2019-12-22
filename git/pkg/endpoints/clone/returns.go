// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cloneReturns, err := UnmarshalCloneReturns(bytes)
//    bytes, err = cloneReturns.Marshal()

package clone

import "encoding/json"

func UnmarshalCloneReturns(data []byte) (CloneReturns, error) {
	var r CloneReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloneReturns struct {
	Location string `json:"location"`
}
