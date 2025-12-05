package response

import (
	"encoding/json"
	"net/http"
)

type APIResponseHandler struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) error {
	response := APIResponseHandler{
		Data:    data,
		Message: message,
		Status:  statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func ErrorResponse(w http.ResponseWriter, statusCode int, er string) error {
	response := APIResponseHandler{
		Message: er,
		Status:  statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}
