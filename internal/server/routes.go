package server

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/superstan777/stock-backend/internal/docs" // import wygenerowanej dokumentacji Swagger
	httpSwagger "github.com/swaggo/http-swagger"

	devicesHandlers "github.com/superstan777/stock-backend/internal/devices/handlers"
	relationsDevicesHandlers "github.com/superstan777/stock-backend/internal/relations/devices/handlers"
	relationsHandlers "github.com/superstan777/stock-backend/internal/relations/handlers"
	relationsUsersHandlers "github.com/superstan777/stock-backend/internal/relations/users/handlers"
	ticketsHandlers "github.com/superstan777/stock-backend/internal/tickets/handlers"
	ticketsStatsHandlers "github.com/superstan777/stock-backend/internal/tickets/stats/handlers"
	usersHandlers "github.com/superstan777/stock-backend/internal/users/handlers"
	worknotesHandlers "github.com/superstan777/stock-backend/internal/worknotes/handlers"
)

func (s *Server) routes() {
	s.Router.Use(LoggingMiddleware)
	s.Router.Use(CORSMiddleware)

	// --- Swagger ---
	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // zmień URL jeśli potrzebne
	))

	s.Router.Route("/api", func(r chi.Router) {

		// --- USERS ---
		r.Route("/users", func(r chi.Router) {
			r.Get("/", usersHandlers.GetUsersHandler)
			r.Get("/{id}", usersHandlers.GetUserHandler)
			r.Post("/", usersHandlers.CreateUserHandler)
			r.Put("/{id}", usersHandlers.UpdateUserHandler)
			r.Delete("/{id}", usersHandlers.DeleteUserHandler)
		})

		// --- DEVICES ---
		r.Route("/devices", func(r chi.Router) {
			r.Get("/", devicesHandlers.GetDevicesHandler)
			r.Get("/{device_type}", devicesHandlers.GetDevicesHandler)
			r.Post("/", devicesHandlers.CreateDeviceHandler)
		})

		r.Route("/device", func(r chi.Router) {
			r.Get("/{id}", devicesHandlers.GetDeviceHandler)
			r.Put("/{id}", devicesHandlers.UpdateDeviceHandler)
			r.Delete("/{id}", devicesHandlers.DeleteDeviceHandler)
		})

		// --- WORKNOTES ---
		r.Route("/worknotes", func(r chi.Router) {
			r.Get("/", worknotesHandlers.GetWorknotesByTicketHandler)
			r.Post("/", worknotesHandlers.CreateWorknoteHandler)
		})

		// --- RELATIONS ---
		r.Route("/relations", func(r chi.Router) {
			r.Post("/", relationsHandlers.CreateRelationHandler)
			r.Patch("/{id}/end", relationsHandlers.EndRelationHandler)
		
			// --- RELATIONS DEVICES ---
			r.Route("/devices/{device_id}/relations", func(r chi.Router) {
				r.Get("/", relationsDevicesHandlers.GetRelationsByDeviceHandler)
				r.Get("/active", relationsDevicesHandlers.HasActiveRelationHandler)
			})

			// --- RELATIONS USERS ---
			r.Route("/users/{user_id}/relations", func(r chi.Router) {
				r.Get("/", relationsUsersHandlers.GetRelationsByUserHandler)
			})
		})

		// --- TICKETS ---
		r.Route("/tickets", func(r chi.Router) {
			r.Get("/", ticketsHandlers.GetTicketsHandler)
			r.Get("/{id}", ticketsHandlers.GetTicketHandler)
			r.Post("/", ticketsHandlers.CreateTicketHandler)
			r.Put("/{id}", ticketsHandlers.UpdateTicketHandler)
			r.Delete("/{id}", ticketsHandlers.DeleteTicketHandler)

			// --- TICKETS STATS ---
			r.Route("/stats", func(r chi.Router) {
				r.Get("/resolved", ticketsStatsHandlers.GetResolvedTicketsStatsHandler)
				r.Get("/open", ticketsStatsHandlers.GetOpenTicketsStatsHandler)
				r.Get("/operators", ticketsStatsHandlers.GetTicketsByOperatorHandler)
			})
		})
	})
}