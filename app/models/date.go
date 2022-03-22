package models

import "time"

type DateAudit struct {
	CreatedAt *time.Time `json:"created_at" example:"2020-07-19T12:27:20.617215Z"`
	UpdatedAt *time.Time `json:"updated_at" example:"2020-07-19T12:27:20.617215Z"`
}
