// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    upsertVolumeReturns, err := UnmarshalUpsertVolumeReturns(bytes)
//    bytes, err = upsertVolumeReturns.Marshal()

package upsertVolume

import "encoding/json"

func UnmarshalUpsertVolumeReturns(data []byte) (UpsertVolumeReturns, error) {
	var r UpsertVolumeReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertVolumeReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UpsertVolumeReturns struct {
	MountPoint string `json:"mountPoint"`
}
