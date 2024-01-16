package dto

import "time"

type Organization struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	DisplayName        string    `json:"display_name"`
	CreateTime         time.Time `json:"create_time"`
	UpdateTime         time.Time `json:"update_time"`
	HasHadValidPayment bool      `json:"has_had_valid_payment"`
}

type OrganizationList struct {
	Items []Organization `json:"items"`
}
