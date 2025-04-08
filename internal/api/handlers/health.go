package handlers

import (
	"net/http"

	"golang_task_manager_folder_structure/internal/logger"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	logger *logger.Logger
}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler(logger *logger.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

// Check performs a health check
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
