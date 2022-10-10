package models

import (
	"github.com/google/uuid"
)

type EspData struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	DeviceId     uuid.UUID `json:"device_id" validate:"required"`
	RandomNumber int       `json:"random_number" validate:"required"`
	CreatedAt    int64     `json:"created_at"`
	UpdatedAt    int64     `json:"updated_at"`
}
