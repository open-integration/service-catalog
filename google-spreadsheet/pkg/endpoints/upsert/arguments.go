// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    upsertArguments, err := UnmarshalUpsertArguments(bytes)
//    bytes, err = upsertArguments.Marshal()

package upsert

import "encoding/json"

func UnmarshalUpsertArguments(data []byte) (UpsertArguments, error) {
	var r UpsertArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UpsertArguments struct {
	Rows           []Row          `json:"Rows"`          // Do upsert
	ServiceAccount ServiceAccount `json:"ServiceAccount"`// service-account
	SpreadsheetID  string         `json:"SpreadsheetID"` // ID of the spreadsheet
}

type Row struct {
	Data []interface{} `json:"Data"`        
	ID   *string       `json:"ID,omitempty"`
}

// service-account
type ServiceAccount struct {
	AuthProviderX509CERTURL string `json:"auth_provider_x509_cert_url"`
	AuthURI                 string `json:"auth_uri"`                   
	ClientEmail             string `json:"client_email"`               
	ClientID                string `json:"client_id"`                  
	ClientX509CERTURL       string `json:"client_x509_cert_url"`       
	PrivateKey              string `json:"private_key"`                
	PrivateKeyID            string `json:"private_key_id"`             
	ProjectID               string `json:"project_id"`                 
	TokenURI                string `json:"token_uri"`                  
	Type                    string `json:"type"`                       
}
