package controllers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/kolesico/FocusGuard/internal/events"
	"github.com/kolesico/FocusGuard/internal/server/response"
)

type EventHandler struct {
	repo events.EventRepository
	log  slog.Logger
}

var validate = validator.New()

func NewEventsHandler(repo events.EventRepository, log slog.Logger) *EventHandler {
	return &EventHandler{repo: repo, log: log}
}

func (e *EventHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /events", e.CreateEvent)
}

func (e *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req events.CreateEventRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		e.log.Error("Invalid request body")
		response.ErrorResponse(w, 501, "Invalid request body")
	}

	err = validate.Struct(req)
	if err != nil {
		e.log.Error("Invalid validate body")
		response.ErrorResponse(w, 501, "Invalid validate body")
	}

	event := &events.Events{
		Event:     req.Event,
		Timestamp: req.Timestamp,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	eventID, err := e.repo.SaveEvent(ctx, event)
	if err != nil {
		e.log.Error("Error save Event in DB", err)
		response.ErrorResponse(w, 501, "Error save Event in DB")
	}

	e.log.Info("Success create event")
	response.SuccessResponse(w, 201, "Success create event", eventID)

}
