package dto

import "time"

type Queue struct {
	ID          string    `json:"id" yaml:"id"`
	Name        string    `json:"name" yaml:"name"`
	DisplayName string    `json:"display_name" yaml:"display_name"`
	CreateTime  time.Time `json:"create_time" yaml:"create_time"`
	UpdateTime  time.Time `json:"update_time" yaml:"update_time"`
}

type QueueList struct {
	Items []Queue `json:"items"`
}
