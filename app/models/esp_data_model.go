package models

import (
	"time"

	"github.com/google/uuid"
)

type EspData struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	DeviceId     uuid.UUID `json:"device_id" validate:"required"`
	RandomNumber int       `json:"random_number" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
