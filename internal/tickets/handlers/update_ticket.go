package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func UpdateTicketHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing ticket ID")
		return
	}

	var input tickets.TicketUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	t, err := repository.Update(db.DB, id, input)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database update error: "+err.Error())
		return
	}

	// Zwracamy spójną odpowiedź dla update
	apiresponse.JSONUpdated(w, t)
}