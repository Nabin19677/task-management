package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"anilkhadka.com.np/task-management/conf"
	"anilkhadka.com.np/task-management/database/postgres"
	"anilkhadka.com.np/task-management/internal/handlers"
	"anilkhadka.com.np/task-management/internal/repositories"
	"anilkhadka.com.np/task-management/utils"

	_ "github.com/lib/pq"
)

// HTMLTemplates contains the parsed HTML templates
var HTMLTemplates *template.Template

// Initialize HTML templates
func initTemplates() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)
}

// StartCronJob starts a cron job for sending daily email reminders
func StartCronJob() {
	// Implementation
}

func main() {
	conf.InitEnvConfigs()

	db := postgres.CreateDBConnection()

	defer db.Close()

	repositories.InitRepositories(db)

	// Initialize HTML templates
	initTemplates()

	// Define HTTP routes handlers
	utils.RegisterRoute("users", handlers.GetUserHandler)
	utils.RegisterRoute("tasks", handlers.GetTaskHandler)

	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	// Start the cron job
	StartCronJob()

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
