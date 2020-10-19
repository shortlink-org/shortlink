package csi_driver

import "fmt"

type accessType int

const (
	mountAccess accessType = iota
	blockAccess
)

type hostPathVolume struct {
	VolName       string     `json:"volName"`
	VolID         string     `json:"volID"`
	VolSize       int64      `json:"volSize"`
	VolPath       string     `json:"volPath"`
	VolAccessType accessType `json:"volAccessType"`
	ParentVolID   string     `json:"parentVolID,omitempty"`
	ParentSnapID  string     `json:"parentSnapID,omitempty"`
	Ephemeral     bool       `json:"ephemeral"`
}

var (
	hostPathVolumes map[string]hostPathVolume
)

func getVolumeByID(volumeID string) (hostPathVolume, error) {
	if hostPathVol, ok := hostPathVolumes[volumeID]; ok {
		return hostPathVol, nil
	}
	return hostPathVolume{}, fmt.Errorf("volume id %s does not exist in the volumes list", volumeID)
}
