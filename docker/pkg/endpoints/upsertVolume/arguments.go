// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    upsertVolumeArguments, err := UnmarshalUpsertVolumeArguments(bytes)
//    bytes, err = upsertVolumeArguments.Marshal()

package upsertVolume

import "encoding/json"

func UnmarshalUpsertVolumeArguments(data []byte) (UpsertVolumeArguments, error) {
	var r UpsertVolumeArguments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertVolumeArguments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UpsertVolumeArguments struct {
	APIVersion string                 `json:"api_version"`     
	Host       string                 `json:"host"`            
	Labels     map[string]interface{} `json:"labels,omitempty"`
	Name       string                 `json:"name"`            
}
