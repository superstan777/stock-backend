package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	var input tickets.TicketInsert

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if input.Title == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Ticket title is required")
		return
	}

	ticket, err := repository.Insert(db.DB, input)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database insert error: "+err.Error())
		return
	}

	// Spójna odpowiedź dla create
	apiresponse.JSONCreated(w, ticket)
}