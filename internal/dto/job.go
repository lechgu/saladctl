package dto

import (
	"encoding/json"
	"time"
)

type JobEvent struct {
	Action string    `json:"action" yaml:"action"`
	Time   time.Time `json:"time" yaml:"time"`
}

type Job struct {
	ID         string          `json:"id" yaml:"id"`
	CreateTime time.Time       `json:"create_time" yaml:"create_time"`
	UpdateTime time.Time       `json:"update_time" yaml:"update_time"`
	Webhook    string          `json:"webhook" yaml:"webhook"`
	Status     string          `json:"status" yaml:"status"`
	Metadata   json.RawMessage `json:"metadata" yaml:"metadata"`
	Input      json.RawMessage `json:"input" yaml:"input"`
	Output     json.RawMessage `json:"output" yaml:"output"`
	Events     []JobEvent      `json:"events" yaml:"events"`
}

type PrettyJob struct {
	ID         string     `yaml:"id"`
	CreateTime time.Time  `yaml:"create_time"`
	UpdateTime time.Time  `yaml:"update_time"`
	Webhook    string     `yaml:"webhook"`
	Status     string     `yaml:"status"`
	Metadata   string     `yaml:"metadata"`
	Input      string     `yaml:"input"`
	Output     string     `yaml:"output"`
	Events     []JobEvent `json:"events" yaml:"events"`
}

func (job Job) MakePretty() PrettyJob {
	return PrettyJob{
		ID:         job.ID,
		CreateTime: job.CreateTime,
		UpdateTime: job.UpdateTime,
		Webhook:    job.Webhook,
		Status:     job.Status,
		Metadata:   string(job.Metadata),
		Input:      string(job.Input),
		Output:     string(job.Output),
		Events:     job.Events,
	}
}
