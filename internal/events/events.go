package events

import (
	"time"
)

type Events struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateEventRequest struct {
	Event     string    `json:"event" validate:"oneof=closed opened"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
}
