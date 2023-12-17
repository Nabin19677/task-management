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

func (s *TaskService) GetByID(taskId int) (*models.Task, error) {
	return s.TaskRepo.GetTask(taskId)
}

func (s *TaskService) CreateTask(newTask *models.NewTask) (bool, error) {
	_, err := s.TaskRepo.Insert(*newTask)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	_, err := s.TaskRepo.Update(*task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetTasksByManager(managerId int) ([]*models.Task, error) {
	tasks, err := s.TaskRepo.GetTasksByManager(managerId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTasksByAssignee(assigneeId int) ([]*models.Task, error) {
	tasks, err := s.TaskRepo.GetTasksByAssignee(assigneeId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) DeleteTask(taskId int) error {
	err := s.TaskRepo.Delete(taskId)
	return err
}
