package services

import (
	"errors"

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

func (s *UserService) CreateUser(newUser *models.NewUser) (bool, error) {
	existingUser, _ := s.UserRepo.FindByEmail(newUser.Email)
	if existingUser != nil {
		return false, errors.New("email already in use")
	}
	_, err := s.UserRepo.Insert(*newUser)
	if err != nil {
		return false, err
	}
	return true, nil
}
