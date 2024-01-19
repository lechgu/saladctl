package dto

import "encoding/json"

type CreateJobRequest struct {
	Input    json.RawMessage `json:"input"`
	Webhook  string          `json:"webhook,omitempty"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}
