package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

func GetComputersHandler(w http.ResponseWriter, r *http.Request) {
	deviceType := "computer"

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	if perPage < 1 {
		perPage = 20
	}

	filters := make(map[string][]string)
	for key, values := range r.URL.Query() {
		if key == "page" || key == "perPage" {
			continue
		}
		var filterValues []string
		for _, v := range values {
			for _, part := range strings.Split(v, ",") {
				part = strings.TrimSpace(part)
				if part != "" {
					filterValues = append(filterValues, part)
				}
			}
		}
		if len(filterValues) > 0 {
			filters[key] = filterValues
		}
	}

	devices, total, err := repository.GetDevicesByType(db.DB, deviceType, filters, page, perPage)
	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data":  devices,
		"count": total,
		"page":  page,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}