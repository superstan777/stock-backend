package apiresponse

import (
	"encoding/json"
	"net/http"
)

// Meta defines pagination and count info.
type Meta struct {
	Count       int `json:"count"`
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
}

// Response defines standard success response with optional metadata.
// swagger:model
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta *Meta       `json:"meta,omitempty"`
}

// PaginatedResponse defines a paginated list response.
// swagger:model
type PaginatedResponse struct {
	// Data contains the list of items
	Data interface{} `json:"data"`
	// Meta contains pagination info
	Meta *Meta `json:"meta"`
}

// ErrorResponse defines error format used across the API.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}

// MessageResponse defines general-purpose message responses.
// swagger:model
type MessageResponse struct {
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// swagger:model
type CreatedResponse struct {
	Message string      `json:"message" example:"Created successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// swagger:model
type UpdatedResponse struct {
	Message string      `json:"message" example:"Updated successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// swagger:model
type DeletedResponse struct {
	Message string            `json:"message" example:"Deleted successfully"`
	Data    map[string]string `json:"data,omitempty" example:"{\"id\":\"123\"}"`
}

// swagger:model
type EndedResponse struct {
	Message string            `json:"message" example:"Relation ended successfully"`
	Data    map[string]string `json:"data,omitempty" example:"{\"id\":\"123\",\"status\":\"ended\"}"`
}

// swagger:model
type ActiveRelationResponse struct {
	Data map[string]bool `json:"data" example:"{\"hasActiveRelation\":true}"`
}

// --- Utility functions ---

func JSONSuccess(w http.ResponseWriter, status int, data interface{}, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Data: data,
		Meta: meta,
	})
}

func JSONPaginated(w http.ResponseWriter, status int, data interface{}, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(PaginatedResponse{
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