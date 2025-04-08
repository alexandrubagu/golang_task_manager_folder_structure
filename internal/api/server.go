package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang_task_manager_folder_structure/internal/api/handlers"
	"golang_task_manager_folder_structure/internal/config"
	"golang_task_manager_folder_structure/internal/logger"
	"golang_task_manager_folder_structure/internal/repository"
	"golang_task_manager_folder_structure/internal/services"
)

// Services contains all service dependencies
type Services struct {
	TaskService *services.TaskService
	Logger      *logger.Logger
}

// NewServices creates a new Services instance
func NewServices(taskRepo *repository.TaskRepository, logger *logger.Logger) *Services {
	return &Services{
		TaskService: services.NewTaskService(taskRepo),
		Logger:      logger,
	}
}

// Server represents the HTTP server
type Server struct {
	router   http.Handler
	server   *http.Server
	services *Services
	logger   *logger.Logger
}

// NewServer creates a new Server instance
func NewServer(cfg *config.Config, services *Services, logger *logger.Logger) *Server {
	server := &Server{
		services: services,
		logger:   logger,
	}

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(services.TaskService, logger)
	healthHandler := handlers.NewHealthHandler(logger)

	// Initialize router
	router := setupRouter(taskHandler, healthHandler, logger)
	server.router = router

	// Configure HTTP server
	server.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server
}

// Start begins listening for requests
func (s *Server) Start() error {
	// Start server in a goroutine
	go func() {
		s.logger.Info("Server starting on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Server failed to start", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	s.logger.Info("Server shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Server forced to shutdown", err)
		return err
	}

	s.logger.Info("Server exited properly")
	return nil
}
