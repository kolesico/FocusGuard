package model

import (
	"time"
)

type Events struct {
	Event     string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
