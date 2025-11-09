package apiresponse

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Count       int `json:"count"`
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta *Meta       `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONSuccess(w http.ResponseWriter, status int, data interface{}, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Data: data,
		Meta: meta,
	})
}

func JSONError(w http.ResponseWriter, status int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: errorMessage,
	})
}

func JSONMessage(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(MessageResponse{
		Message: message,
		Data:    data,
	})
}

func JSONCreated(w http.ResponseWriter, data interface{}) {
	JSONMessage(w, http.StatusCreated, "Created successfully", data)
}

func JSONUpdated(w http.ResponseWriter, data interface{}) {
	JSONMessage(w, http.StatusOK, "Updated successfully", data)
}

func JSONDeleted(w http.ResponseWriter, id string) {
	JSONMessage(w, http.StatusOK, "Deleted successfully", map[string]string{"id": id})
}

func JSONEnded(w http.ResponseWriter, id string) {
	JSONMessage(w, http.StatusOK, "Relation ended successfully", map[string]string{
		"id":     id,
		"status": "ended",
	})
}

func JSONHasActiveRelation(w http.ResponseWriter, status int, active bool) {
	JSONSuccess(w, status, map[string]bool{"hasActiveRelation": active}, nil)
}