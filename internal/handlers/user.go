package handlers

import (
	"net/http"

	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/utils"
)

// Implement HTTP handlers for task management
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	userID := utils.ExtractIDFromURL(r.URL.Path, "/users/")
	if userID != "" {
		userService := services.NewUserService()
		userService.GetByID(1)

	} else {

	}
}

func handlePOST(w http.ResponseWriter, r *http.Request) {

}
