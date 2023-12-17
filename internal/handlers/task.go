package handlers

import (
	"log"
	"net/http"
	"strconv"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

// Implement HTTP handlers for task management
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskService := services.NewTaskService()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")
		dueDate := r.FormValue("dueDate")
		status := r.FormValue("status")

		managerId, err := strconv.Atoi(r.FormValue("manager"))
		if err != nil {
			return
		}

		assigneeId, err := strconv.Atoi(r.FormValue("assignee"))
		if err != nil {
			return
		}

		newTask := &models.NewTask{
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			Status:      status,
			ManagerID:   managerId,
			AssigneeID:  assigneeId,
		}

		success, err := taskService.CreateTask(newTask)
		if err != nil {
			log.Println(err)
		}

		if success {
			taskService := services.NewTaskService()

			tasks, _ := taskService.GetTasksByManager(1)

			pageVariables := types.PageVariables{
				Title: "Dashboard",
				Data: map[string]interface{}{
					"Tasks": tasks,
				},
			}
			utils.RenderTemplate(w, "dashboard.html", pageVariables)
		} else {
			return
		}

	} else if r.Method == http.MethodGet {
		pageVariables := types.PageVariables{
			Title: "Create Task",
			Data:  nil,
		}
		utils.RenderTemplate(w, "create_task.html", pageVariables)
	}
}
