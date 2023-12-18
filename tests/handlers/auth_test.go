package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"anilkhadka.com.np/task-management/internal/handlers"
)

func TestLogoutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handlers.LogoutHandler(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	expectedURL := "/login"
	if location := rr.Header().Get("Location"); location != expectedURL {
		t.Errorf("Handler returned unexpected Location header: got %v want %v",
			location, expectedURL)
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected one cookie, but got %d", len(cookies))
	} else {
		cookie := cookies[0]
		if cookie.Name != "auth_token" {
			t.Errorf("Expected cookie with name 'auth_token'")
		}
		if cookie.Value != "" {
			t.Errorf("Expected cookie value to be empty, but got %s", cookie.Value)
		}
	}
}

// func TestLoginHandler(t *testing.T) {
// 	mockUserService := new(mocks.MockUserService)

// 	req, err := http.NewRequest("POST", "/login", strings.NewReader("username=test&password=test"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	mockUserService.On("LoginUser", mock.AnythingOfType("*models.LoginInput")).Return(&models.AuthResponse{}, nil)

// 	handlers.LoginHandler(rr, req)

// 	assert.Equal(t, http.StatusSeeOther, rr.Code)
// }
