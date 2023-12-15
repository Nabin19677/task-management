package services

import (
	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
	TaskRepo *repositories.TaskRepository
}

func NewUserService() *UserService {
	repository := repositories.GetRepository()
	return &UserService{
		UserRepo: repository.User,
		TaskRepo: repository.Task,
	}
}

func (s *UserService) GetByID(userID int) (*models.User, error) {
	return s.UserRepo.GetUserByID(userID)
}
