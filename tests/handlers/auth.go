package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"anilkhadka.com.np/task-management/internal/handlers"
	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginHandler(t *testing.T) {
	mockUserService := new(mocks.MockUserService)

	req, err := http.NewRequest("POST", "/login", strings.NewReader("username=test&password=test"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mockUserService.On("LoginUser", mock.AnythingOfType("*models.LoginInput")).Return(&models.AuthResponse{}, nil)

	handlers.LoginHandler(rr, req)

	assert.Equal(t, http.StatusSeeOther, rr.Code)
}
