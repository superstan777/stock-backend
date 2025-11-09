package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing ticket ID")
		return
	}

	if err := repository.Delete(db.DB, id); err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database delete error: "+err.Error())
		return
	}

	// Spójna odpowiedź dla delete
	apiresponse.JSONDeleted(w, id)
}