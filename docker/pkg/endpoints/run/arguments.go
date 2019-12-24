// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    runArguments, err := UnmarshalRunArguments(bytes)
//    bytes, err = runArguments.Marshal()

package run

import "encoding/json"

func UnmarshalRunArguments(data []byte) (RunArguments, error) {
	var r RunArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RunArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RunArguments struct {
	APIVersion           string   `json:"api_version"`         
	Arguments            []string `json:"arguments"`           
	Command              string   `json:"command"`             
	EnvironmentVariables []string `json:"environmentVariables"`
	Host                 string   `json:"host"`                
	Image                string   `json:"image"`               
	Volumes              []string `json:"volumes"`             
}
