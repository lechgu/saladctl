package dto

type CreateQueueRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"dispay_name"`
	Description string `json:"description"`
}
