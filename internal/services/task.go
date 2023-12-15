package services

import (
	"database/sql"

	"anilkhadka.com.np/task-management/internal/models"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(taskID int) error
	AssignTask(taskID int, assignee string) error
	MarkTaskAsDone(taskID int) error
	GetAllTasks() ([]*models.Task, error)
}

// Service implements TaskService interface
type Service struct {
	DB *sql.DB
}

// Implement TaskService methods for Service

func (m *Service) CreateTask(task *models.Task) error {
	// Implementation
	return nil
}

func (m *Service) UpdateTask(task *models.Task) error {
	// Implementation
	return nil
}

func (m *Service) DeleteTask(taskID int) error {
	// Implementation
	return nil
}

func (m *Service) AssignTask(taskID int, assignee string) error {
	// Implementation
	return nil
}

func (m *Service) MarkTaskAsDone(taskID int) error {
	// Implementation
	return nil
}

func (m *Service) GetAllTasks() ([]*models.Task, error) {
	// Implementation
	return nil, nil
}
