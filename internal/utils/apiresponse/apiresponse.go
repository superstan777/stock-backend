package apiresponse

import (
	"encoding/json"
	"net/http"
)

// Meta zawiera informacje o liczbie wyników i ewentualnie paginacji.
type Meta struct {
	Count       int `json:"count"`
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
}

// Response to ujednolicona struktura odpowiedzi API (z danymi i metadanymi).
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta *Meta       `json:"meta,omitempty"`
}

// ErrorResponse to ujednolicona struktura błędu.
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse to uproszczona struktura odpowiedzi z komunikatem.
type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSONSuccess tworzy standardową odpowiedź z danymi i metadanymi (paginacja lub nie).
func JSONSuccess(w http.ResponseWriter, status int, data interface{}, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Data: data,
		Meta: meta,
	})
}

// JSONError zwraca błąd API.
func JSONError(w http.ResponseWriter, status int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: errorMessage,
	})
}

// JSONMessage zwraca prosty komunikat sukcesu dla create, update, delete
func JSONMessage(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(MessageResponse{
		Message: message,
		Data:    data,
	})
}

// Spójne helpery dla create, update, delete
// Spójne helpery dla create, update, delete, end
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

// JSONHasActiveRelation zwraca dedykowany obiekt dla endpointu hasActiveRelation
// data.hasActiveRelation: true/false
func JSONHasActiveRelation(w http.ResponseWriter, status int, active bool) {
	JSONSuccess(w, status, map[string]bool{"hasActiveRelation": active}, nil)
}