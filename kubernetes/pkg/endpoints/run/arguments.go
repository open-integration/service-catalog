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
	Auth     Auth     `json:"Auth"`              
	Detached *bool    `json:"Detached,omitempty"`// Start the pod and do not wait for it(ignoring any timeout argument)
	Pod      string   `json:"Pod"`               
	Timeout  *float64 `json:"Timeout,omitempty"` // How long to wait for the pod to finished before sending termination request(when detached; argument provided, this property ignored)
}

type Auth struct {
	CRT   *string `json:"Crt,omitempty"`  
	Host  *string `json:"Host,omitempty"` 
	Token *string `json:"Token,omitempty"`
	Type  Type    `json:"Type"`           
}

type Type string
const (
	KubernetesServiceAccount Type = "KubernetesServiceAccount"
)
