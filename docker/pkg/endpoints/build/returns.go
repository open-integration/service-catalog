// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    buildReturns, err := UnmarshalBuildReturns(bytes)
//    bytes, err = buildReturns.Marshal()

package build

import "encoding/json"

type BuildReturns map[string]interface{}

func UnmarshalBuildReturns(data []byte) (BuildReturns, error) {
	var r BuildReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuildReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
