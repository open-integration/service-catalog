// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    commandReturns, err := UnmarshalCommandReturns(bytes)
//    bytes, err = commandReturns.Marshal()

package command

import "encoding/json"

type CommandReturns map[string]interface{}

func UnmarshalCommandReturns(data []byte) (CommandReturns, error) {
	var r CommandReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CommandReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
