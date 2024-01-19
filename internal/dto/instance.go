package dto

import "time"

type Instance struct {
	MachineID  string    `json:"machine_id" yaml:"machine_id"`
	State      string    `json:"state" yaml:"state"`
	UpdateTime time.Time `json:"update_time" yaml:"update_time"`
	Version    int       `json:"version" yaml:"version"`
}

type InstanceList struct {
	Instances []Instance `json:"instances"`
}
