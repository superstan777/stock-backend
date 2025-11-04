package server

import (
	"github.com/go-chi/chi/v5"

	userHandlers "github.com/superstan777/stock-backend/internal/users/handlers"
	// deviceHandlers "github.com/superstan777/stock-backend/internal/devices/handlers"
)

func (s *Server) routes() {
	s.Router.Use(LoggingMiddleware)
	s.Router.Use(CORSMiddleware)

	s.Router.Route("/api", func(r chi.Router) {

		// --- USERS ---
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandlers.GetUsersHandler)
			r.Get("/{id}", userHandlers.GetUserHandler)
			r.Post("/", userHandlers.AddUserHandler)
			r.Put("/{id}", userHandlers.UpdateUserHandler)
			r.Delete("/{id}", userHandlers.DeleteUserHandler)
		})

		// --- DEVICES (zakomentowane do czasu utworzenia modułu) ---
		/*
		r.Route("/devices", func(r chi.Router) {
			r.Get("/", deviceHandlers.GetDevicesHandler)
			r.Get("/{id}", deviceHandlers.GetDeviceHandler)
			r.Post("/", deviceHandlers.AddDeviceHandler)
			r.Put("/{id}", deviceHandlers.UpdateDeviceHandler)
			r.Delete("/{id}", deviceHandlers.DeleteDeviceHandler)
		})
		*/
	})
}