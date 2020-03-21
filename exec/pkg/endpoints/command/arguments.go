// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    commandArguments, err := UnmarshalCommandArguments(bytes)
//    bytes, err = commandArguments.Marshal()

package command

import "encoding/json"

func UnmarshalCommandArguments(data []byte) (CommandArguments, error) {
	var r CommandArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CommandArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CommandArguments struct {
	Command string `json:"Command"`
}
