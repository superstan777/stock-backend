package handlers

import (
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/users/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page := 1
	perPage := 20
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if p := query.Get("per_page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			perPage = parsed
		}
	}

	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" || key == "per_page" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	usersList, count, err := repository.GetUsers(db.DB, filters, page, perPage)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	meta := &apiresponse.Meta{
		Count:       count,
		CurrentPage: page,
		TotalPages:  (count + perPage - 1) / perPage,
		PerPage:     perPage,
	}

	apiresponse.JSONSuccess(w, http.StatusOK, usersList, meta)
}