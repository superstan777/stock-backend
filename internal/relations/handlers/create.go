package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations"
	"github.com/superstan777/stock-backend/internal/relations/repository"
)

func CreateRelationHandler(w http.ResponseWriter, r *http.Request) {
	var input relations.RelationInsert
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// basic validation
	if input.DeviceID == "" || input.UserID == "" {
		http.Error(w, "missing device_id or user_id", http.StatusBadRequest)
		return
	}

	// if start date not provided, use now
	if input.StartDate.IsZero() {
		input.StartDate = time.Now().UTC()
	}

	rel, err := repository.Insert(db.DB, input)
	if err != nil {
		http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rel)
}