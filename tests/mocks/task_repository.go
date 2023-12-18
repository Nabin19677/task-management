package mocks

import (
	"anilkhadka.com.np/task-management/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository is a mock implementation of the TaskRepository interface.
type MockTaskRepository struct {
	mock.Mock
}

// Insert is a mocked implementation of the Insert method in the TaskRepository interface.
func (m *MockTaskRepository) Insert(newTask models.NewTask) (*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(newTask)

	// Extract and return the mocked results
	return args.Get(0).(*models.Task), args.Error(1)
}

// GetTask is a mocked implementation of the GetTask method in the TaskRepository interface.
func (m *MockTaskRepository) GetTask(taskID int) (*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(taskID)

	// Extract and return the mocked results
	return args.Get(0).(*models.Task), args.Error(1)
}

// Update is a mocked implementation of the Update method in the TaskRepository interface.
func (m *MockTaskRepository) Update(task models.Task) (*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(task)

	// Extract and return the mocked results
	return args.Get(0).(*models.Task), args.Error(1)
}

// UpdateTaskStatus is a mocked implementation of the UpdateTaskStatus method in the TaskRepository interface.
func (m *MockTaskRepository) UpdateTaskStatus(taskID int, newStatus string) (*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(taskID, newStatus)

	// Extract and return the mocked results
	return args.Get(0).(*models.Task), args.Error(1)
}

// GetTasksByManager is a mocked implementation of the GetTasksByManager method in the TaskRepository interface.
func (m *MockTaskRepository) GetTasksByManager(managerID int) ([]*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(managerID)

	// Extract and return the mocked results
	return args.Get(0).([]*models.Task), args.Error(1)
}

// GetTasksByAssignee is a mocked implementation of the GetTasksByAssignee method in the TaskRepository interface.
func (m *MockTaskRepository) GetTasksByAssignee(assigneeID int) ([]*models.Task, error) {
	// Arguments that were passed during the method call
	args := m.Called(assigneeID)

	// Extract and return the mocked results
	return args.Get(0).([]*models.Task), args.Error(1)
}

// Delete is a mocked implementation of the Delete method in the TaskRepository interface.
func (m *MockTaskRepository) Delete(taskID int) error {
	// Arguments that were passed during the method call
	args := m.Called(taskID)

	// Extract and return the mocked results
	return args.Error(0)
}

// AssigneeWithUndoneTasks is a mocked implementation of the AssigneeWithUndoneTasks method in the TaskRepository interface.
func (m *MockTaskRepository) AssigneeWithUndoneTasks() ([]models.AssigneeWithTask, error) {
	// Arguments that were passed during the method call
	args := m.Called()

	// Extract and return the mocked results
	return args.Get(0).([]models.AssigneeWithTask), args.Error(1)
}
