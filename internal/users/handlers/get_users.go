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
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	usersList, count, err := repository.GetUsers(db.DB, filters, page, 20)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	meta := map[string]interface{}{
		"count":        count,
		"current_page": page,
		"total_pages":  (count + 20 - 1) / 20, 
	}

	apiresponse.JSONSuccess(w, http.StatusOK, "Users fetched successfully", map[string]interface{}{
		"users": usersList,
		"meta":  meta,
	})
}