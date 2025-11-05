package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/worknotes/repository"
)

// GET /worknotes?ticket_id={id}
func GetWorknotesByTicketHandler(w http.ResponseWriter, r *http.Request) {
	ticketID := r.URL.Query().Get("ticket_id")
	if ticketID == "" {
		http.Error(w, "missing ticket_id parameter", http.StatusBadRequest)
		return
	}

	notes, err := repository.GetWorknotesByTicket(db.DB, ticketID)
	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}