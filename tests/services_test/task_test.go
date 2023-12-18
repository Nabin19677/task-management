package services

import (
	"testing"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockRepo := &mocks.MockTaskRepository{}

	taskService := services.NewTaskService()

	newTask := &models.NewTask{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	expectedTask := &models.Task{
		ID:          1,
		Title:       newTask.Title,
		Description: newTask.Description,
	}

	mockRepo.On("Insert", *newTask).Return(expectedTask, nil)

	result, err := taskService.CreateTask(newTask)

	assert.NoError(t, err, "CreateTask should not return an error")
	assert.True(t, result, "CreateTask should return true for success")

	mockRepo.AssertExpectations(t)
}
