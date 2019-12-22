// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cloneArguments, err := UnmarshalCloneArguments(bytes)
//    bytes, err = cloneArguments.Marshal()

package clone

import "encoding/json"

func UnmarshalCloneArguments(data []byte) (CloneArguments, error) {
	var r CloneArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloneArguments struct {
	Auth     *GithubAuth `json:"auth,omitempty"`
	Path     string      `json:"path"`          
	Provider Provider    `json:"provider"`      
	Repo     string      `json:"repo"`          
}

type GithubAuth struct {
	Token    *string `json:"token,omitempty"`   
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

type Provider string
const (
	Github Provider = "github"
)
