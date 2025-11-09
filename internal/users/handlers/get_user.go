package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/users/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	user, err := repository.GetByID(db.DB, id)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	if user == nil {
		apiresponse.JSONError(w, http.StatusNotFound, "User not found")
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, user, nil)
}