package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang_task_manager_folder_structure/internal/logger"
	"golang_task_manager_folder_structure/internal/services"

	"github.com/go-chi/chi/v5"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service *services.TaskService
	logger  *logger.Logger
}

// TaskRequest represents a task request body
type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date,omitempty"`
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(service *services.TaskService, logger *logger.Logger) *TaskHandler {
	return &TaskHandler{
		service: service,
		logger:  logger,
	}
}

// List returns all tasks
func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll()
	if err != nil {
		h.logger.Error("Failed to get tasks", err)
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	respondJSON(w, tasks, http.StatusOK)
}

// Get returns a specific task
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get task", err)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	respondJSON(w, task, http.StatusOK)
}

// Create adds a new task
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.Create(req.Title, req.Description, req.DueDate)
	if err != nil {
		h.logger.Error("Failed to create task", err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, task, http.StatusCreated)
}

// Update modifies an existing task
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.Update(id, req.Title, req.Description, req.DueDate)
	if err != nil {
		h.logger.Error("Failed to update task", err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, task, http.StatusOK)
}

// Delete removes a task
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		h.logger.Error("Failed to delete task", err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Complete marks a task as completed
func (h *TaskHandler) Complete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.Complete(id)
	if err != nil {
		h.logger.Error("Failed to complete task", err)
		http.Error(w, "Failed to complete task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, task, http.StatusOK)
}

// Helper function to respond with JSON
func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
