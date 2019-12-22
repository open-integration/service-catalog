// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    buildArguments, err := UnmarshalBuildArguments(bytes)
//    bytes, err = buildArguments.Marshal()

package build

import "encoding/json"

func UnmarshalBuildArguments(data []byte) (BuildArguments, error) {
	var r BuildArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuildArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BuildArguments struct {
	APIVersion   string `json:"api_version"`  
	BuildContext string `json:"build_context"`
	Dockerfile   string `json:"dockerfile"`   
	Host         string `json:"host"`         
	Path         string `json:"path"`         
	Tag          string `json:"tag"`          
}
