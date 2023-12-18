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
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	taskService := services.NewTaskService()
	userService := services.NewUserService()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")
		dueDate := r.FormValue("dueDate")

		managerId, err := strconv.Atoi(userID)
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
			Status:      types.StatusTodo,
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
			utils.RenderTemplate(w, "manager_dashboard.html", pageVariables)
		} else {
			return
		}

	} else if r.Method == http.MethodGet {
		assignees, err := userService.FindUsersByRole(types.Assignee)

		if err != nil {
			log.Println(err)
		}
		pageVariables := types.PageVariables{
			Title: "Create Task",
			Data: map[string]interface{}{
				"Assignees": assignees,
			},
		}
		utils.RenderTemplate(w, "create_task.html", pageVariables)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(r.URL.Path[len("/delete-task/"):])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	taskService := services.NewTaskService()

	err = taskService.DeleteTask(taskID)

	if err != nil {
		log.Fatal(err)
		return

	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the URL
	taskID, err := strconv.Atoi(r.URL.Path[len("/edit-task/"):])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	taskService := services.NewTaskService()
	if r.Method == http.MethodPost {
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get the form values
		title := r.FormValue("title")
		description := r.FormValue("description")
		dueDate := r.FormValue("dueDate")
		status := r.FormValue("status")

		// Create a new task object with the updated values
		updatedTask := &models.Task{
			ID:          taskID,
			Title:       title,
			Description: description,
			DueDate:     dueDate,
			Status:      types.TaskStatus(status),

			// Set other fields accordingly
		}

		err = taskService.UpdateTask(updatedTask)
		if err != nil {
			http.Error(w, "Failed to update task", http.StatusInternalServerError)
			return
		}

		// Redirect the user to the dashboard or task list page after successful update
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		// Fetch the task details from the database based on the task ID
		task, err := taskService.GetByID(taskID)

		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to fetch task details", http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to fetch assignees", http.StatusInternalServerError)
			return
		}

		// Render the edit task page with the task details and additional data
		pageVariables := types.PageVariables{
			Title: "Edit Task",
			Data: map[string]interface{}{
				"Task": task,
			},
		}
		utils.RenderTemplate(w, "edit_task.html", pageVariables)
	}

}

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from the URL
	taskID, err := strconv.Atoi(r.URL.Path[len("/update-status/"):])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Parse the form values
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	// Get the new status from the form
	newStatus := r.FormValue("status")

	// Update the task status in the repository
	taskService := services.NewTaskService()
	err = taskService.UpdateTaskStatus(taskID, newStatus)
	if err != nil {
		http.Error(w, "Failed to update task status", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
