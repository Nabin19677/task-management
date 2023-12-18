package mocks

import (
	"anilkhadka.com.np/task-management/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) LoginUser(loginInput *models.LoginInput) (*models.AuthResponse, error) {
	args := m.Called(loginInput)
	return args.Get(0).(*models.AuthResponse), args.Error(1)
}
