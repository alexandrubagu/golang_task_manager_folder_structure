package services

import (
	"errors"
	"time"

	"golang_task_manager_folder_structure/internal/repository"
)

// TaskService handles business logic for tasks
type TaskService struct {
	repo *repository.TaskRepository
}

// NewTaskService creates a new TaskService
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// GetAll returns all tasks
func (s *TaskService) GetAll() ([]repository.Task, error) {
	return s.repo.FindAll()
}

// GetByID returns a task by ID
func (s *TaskService) GetByID(id int) (*repository.Task, error) {
	return s.repo.FindByID(id)
}

// Create adds a new task
func (s *TaskService) Create(title, description, dueDate string) (*repository.Task, error) {
	var due *time.Time

	if dueDate != "" {
		parsedDate, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return nil, errors.New("invalid due date format, expected YYYY-MM-DD")
		}
		due = &parsedDate
	}

	task := &repository.Task{
		Title:       title,
		Description: description,
		DueDate:     due,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Create(task)
}

// Update modifies an existing task
func (s *TaskService) Update(id int, title, description, dueDate string) (*repository.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		task.Title = title
	}

	if description != "" {
		task.Description = description
	}

	if dueDate != "" {
		parsedDate, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return nil, errors.New("invalid due date format, expected YYYY-MM-DD")
		}
		task.DueDate = &parsedDate
	}

	task.UpdatedAt = time.Now()

	return s.repo.Update(task)
}

// Delete removes a task
func (s *TaskService) Delete(id int) error {
	return s.repo.Delete(id)
}

// Complete marks a task as completed
func (s *TaskService) Complete(id int) (*repository.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	task.Completed = true
	task.CompletedAt = func() *time.Time { now := time.Now(); return &now }()
	task.UpdatedAt = time.Now()

	return s.repo.Update(task)
}
