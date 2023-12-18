package services

import (
	"errors"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/repositories"
	"anilkhadka.com.np/task-management/internal/types"
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

func (s *UserService) FindUsersByRole(role types.Role) ([]*models.PublicUser, error) {
	users, err := s.UserRepo.FindByRole(role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) LoginUser(loginInput *models.LoginInput) (*models.AuthResponse, error) {
	user, err := s.UserRepo.FindByEmail(loginInput.Email)
	if err != nil {
		return nil, errors.New("email or password is wrongss")
	}

	err = user.ComparePassword(loginInput.Password)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	token, err := user.GenToken()

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
	}, nil
}
