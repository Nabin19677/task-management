package services

import (
	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/repositories"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

func NewTaskService() *TaskService {
	repository := repositories.GetRepository()
	return &TaskService{
		TaskRepo: repository.Task,
	}
}

func (s *TaskService) CreateTask(newTask *models.NewTask) (bool, error) {
	_, err := s.TaskRepo.Insert(*newTask)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *TaskService) GetTasksByManager(managerId int) ([]*models.Task, error) {
	tasks, err := s.TaskRepo.GetTasksByManager(managerId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
