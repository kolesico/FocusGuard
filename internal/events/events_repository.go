package events

import (
	"context"
)

type EventRepository interface {
	SaveEvent(context.Context, *Events) (int64, error)
}
