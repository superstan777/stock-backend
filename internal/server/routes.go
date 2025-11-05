package server

import (
	"github.com/go-chi/chi/v5"

	deviceHandlers "github.com/superstan777/stock-backend/internal/devices/handlers"
	userHandlers "github.com/superstan777/stock-backend/internal/users/handlers"
	worknotesHandlers "github.com/superstan777/stock-backend/internal/worknotes/handlers"
)

func (s *Server) routes() {
	s.Router.Use(LoggingMiddleware)
	s.Router.Use(CORSMiddleware)

	s.Router.Route("/api", func(r chi.Router) {


		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandlers.GetUsersHandler)
			r.Get("/{id}", userHandlers.GetUserHandler)
			r.Post("/", userHandlers.AddUserHandler)
			r.Put("/{id}", userHandlers.UpdateUserHandler)
			r.Delete("/{id}", userHandlers.DeleteUserHandler)
		})

		r.Route("/devices", func(r chi.Router) {
			r.Get("/computers", deviceHandlers.GetComputersHandler)
   			r.Get("/monitors", deviceHandlers.GetMonitorsHandler)
   			r.Get("/", deviceHandlers.GetAllDevicesHandler) 
			r.Post("/", deviceHandlers.CreateDeviceHandler)
			r.Get("/{id}", deviceHandlers.GetDeviceHandler)
			r.Put("/{id}", deviceHandlers.UpdateDeviceHandler)
			r.Delete("/{id}", deviceHandlers.DeleteDeviceHandler)
		})

		r.Route("/worknotes", func(r chi.Router) {
			r.Get("/", worknotesHandlers.GetWorknotesByTicketHandler)
			r.Post("/", worknotesHandlers.CreateWorknoteHandler)
		})
	
	})
}