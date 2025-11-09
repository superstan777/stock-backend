package handlers

import (
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
	"github.com/superstan777/stock-backend/internal/worknotes/repository"
)

func GetWorknotesByTicketHandler(w http.ResponseWriter, r *http.Request) {
	ticketID := r.URL.Query().Get("ticket_id")
	if ticketID == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing ticket_id parameter")
		return
	}

	notes, err := repository.GetWorknotesByTicket(db.DB, ticketID)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	// Zwracamy dane jako lista, meta nil bo brak paginacji
	apiresponse.JSONSuccess(w, http.StatusOK, notes, nil)
}