package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/users"
	"github.com/superstan777/stock-backend/internal/users/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input users.UserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if input.Name == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if input.Email == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Email is required")
		return
	}

	user, err := repository.Insert(db.DB, input)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database insert error: "+err.Error())
		return
	}

	apiresponse.JSONSuccess(w, http.StatusCreated, "User created successfully", user)
}