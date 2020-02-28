// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    createpvcArguments, err := UnmarshalCreatepvcArguments(bytes)
//    bytes, err = createpvcArguments.Marshal()

package createpvc

import "encoding/json"

func UnmarshalCreatepvcArguments(data []byte) (CreatepvcArguments, error) {
	var r CreatepvcArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatepvcArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreatepvcArguments struct {
	Auth Auth   `json:"Auth"`
	Pvc  string `json:"PVC"` 
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
