package main

import (
	"fmt"
	"log"
	"net/http"

	"anilkhadka.com.np/task-management/conf"
	"anilkhadka.com.np/task-management/database/postgres"
	"anilkhadka.com.np/task-management/internal/handlers"
	"anilkhadka.com.np/task-management/internal/middlewares"
	"anilkhadka.com.np/task-management/internal/repositories"
	"anilkhadka.com.np/task-management/utils"

	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

// Initialize HTML templates
func initTemplates() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.Handle("/create-task", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreateTaskHandler)))
	http.Handle("/delete-task/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeleteTaskHandler)))
	http.Handle("/dashboard", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DashboardHandler)))
}

func StartCronJobs() {
	c := cron.New()

	_, err := c.AddFunc("0 9 * * *", handlers.SendDailyMail)

	if err != nil {
		log.Fatal(err)
	}

	c.Start()
}

func main() {
	conf.InitEnvConfigs()

	db := postgres.CreateDBConnection()

	defer db.Close()

	repositories.InitRepositories(db)

	// Initialize HTML templates
	initTemplates()

	// Define HTTP API routes handlers
	utils.RegisterRoute("users", handlers.GetUserHandler)
	utils.RegisterRoute("tasks", handlers.GetTaskHandler)

	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	// Start the cron job
	StartCronJobs()

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
