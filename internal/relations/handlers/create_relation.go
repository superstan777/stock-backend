package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations"
	"github.com/superstan777/stock-backend/internal/relations/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func CreateRelationHandler(w http.ResponseWriter, r *http.Request) {
	var input relations.RelationInsert
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if input.DeviceID == "" || input.UserID == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device_id or user_id")
		return
	}

	if input.StartDate.IsZero() {
		input.StartDate = time.Now().UTC()
	}

	rel, err := repository.Insert(db.DB, input)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database insert error: "+err.Error())
		return
	}

	apiresponse.JSONSuccess(w, http.StatusCreated, "Relation created successfully", rel)
}