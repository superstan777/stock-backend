package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
)

func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	if err := repository.Delete(db.DB, id); err != nil {
		http.Error(w, "DB delete error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}