package api

import (
	"golang_task_manager_folder_structure/internal/api/handlers"
	"golang_task_manager_folder_structure/internal/api/middlewares"
	"golang_task_manager_folder_structure/internal/logger"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// setupRouter configures the router with all routes and middlewares
func setupRouter(taskHandler *handlers.TaskHandler, healthHandler *handlers.HealthHandler, logger *logger.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middlewares.LoggerMiddleware(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Get("/health", healthHandler.Check)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", taskHandler.List)
			r.Post("/", taskHandler.Create)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", taskHandler.Get)
				r.Put("/", taskHandler.Update)
				r.Delete("/", taskHandler.Delete)
				r.Put("/complete", taskHandler.Complete)
			})
		})
	})

	return r
}
