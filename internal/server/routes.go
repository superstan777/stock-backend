package server

import (
	"github.com/go-chi/chi/v5"

	deviceHandlers "github.com/superstan777/stock-backend/internal/devices/handlers"
	relationHandlers "github.com/superstan777/stock-backend/internal/relations/handlers"
	ticketHandlers "github.com/superstan777/stock-backend/internal/tickets/handlers"
	userHandlers "github.com/superstan777/stock-backend/internal/users/handlers"
	worknotesHandlers "github.com/superstan777/stock-backend/internal/worknotes/handlers"
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

		// --- DEVICES ---
		r.Route("/devices", func(r chi.Router) {
			r.Get("/", deviceHandlers.GetDevicesHandler)
    		r.Get("/{device_type}", deviceHandlers.GetDevicesHandler) 
    		r.Post("/", deviceHandlers.CreateDeviceHandler)       
		})

		r.Route("/device", func(r chi.Router) {
    		r.Get("/{id}", deviceHandlers.GetDeviceHandler)
    		r.Put("/{id}", deviceHandlers.UpdateDeviceHandler)
    		r.Delete("/{id}", deviceHandlers.DeleteDeviceHandler)
		})


		// --- WORKNOTES ---
		r.Route("/worknotes", func(r chi.Router) {
			r.Get("/", worknotesHandlers.GetWorknotesByTicketHandler)
			r.Post("/", worknotesHandlers.CreateWorknoteHandler)
		})

		// --- RELATIONS ---
		r.Route("/relations", func(r chi.Router) {
			r.Get("/device/{device_id}", relationHandlers.GetRelationsByDeviceHandler)
			r.Get("/user/{user_id}", relationHandlers.GetRelationsByUserHandler)
			r.Get("/device/{device_id}/active", relationHandlers.HasActiveRelationHandler) 
			r.Post("/", relationHandlers.CreateRelationHandler)
			r.Post("/{id}/end", relationHandlers.EndRelationHandler)
		})

		// --- TICKETS ---
		r.Route("/tickets", func(r chi.Router) {
			r.Get("/", ticketHandlers.GetTicketsHandler)
			r.Get("/{id}", ticketHandlers.GetTicketHandler)
			r.Post("/", ticketHandlers.AddTicketHandler)
			r.Put("/{id}", ticketHandlers.UpdateTicketHandler)
			r.Delete("/{id}", ticketHandlers.DeleteTicketHandler)
		})
	})
}