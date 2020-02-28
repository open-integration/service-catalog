// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    createpvcReturns, err := UnmarshalCreatepvcReturns(bytes)
//    bytes, err = createpvcReturns.Marshal()

package createpvc

import "encoding/json"

type CreatepvcReturns map[string]interface{}

func UnmarshalCreatepvcReturns(data []byte) (CreatepvcReturns, error) {
	var r CreatepvcReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatepvcReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
