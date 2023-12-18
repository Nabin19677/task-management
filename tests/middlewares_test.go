package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"anilkhadka.com.np/task-management/internal/middlewares"
)

func TestAuthMiddleware(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name           string
		cookieValue    string
		expectedStatus int
	}{
		{"ValidToken", "valid_token_value", http.StatusOK},
		{"InvalidToken", "", http.StatusSeeOther},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.cookieValue != "" {
				req.AddCookie(&http.Cookie{Name: "auth_token", Value: tt.cookieValue})
			}

			rr := httptest.NewRecorder()

			middlewares.AuthMiddleware(dummyHandler).ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Unexpected status code. Expected %d, got %d.", tt.expectedStatus, status)
			}

		})
	}
}
