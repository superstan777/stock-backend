package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// convenience helper: wysyłanie zwykłego JSON-a z layoutem
func WriteJSON(w http.ResponseWriter, status int, data interface{}, meta interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIResponse{
		Success: status >= 200 && status < 300,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	// dodaj timestamp do meta jeśli meta == nil? można, ale nie musimy
	_ = json.NewEncoder(w).Encode(resp)
}

// convenience helper: wysyłanie błędu w ustandaryzowanym formacie
func WriteError(w http.ResponseWriter, status int, message string, errDetail interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := APIResponse{
		Success: false,
		Message: message,
		Error: map[string]interface{}{
			"detail":   errDetail,
			"status":   status,
			"ts":       time.Now().UTC().Format(time.RFC3339),
		},
	}
	_ = json.NewEncoder(w).Encode(resp)
}