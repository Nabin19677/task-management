package handlers

import (
	"net/http"

	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	taskService := services.NewTaskService()

	tasks, _ := taskService.GetTasksByManager(1)

	pageVariables := types.PageVariables{
		Title: "Dashboard",
		Data: map[string]interface{}{
			"Tasks": tasks,
		},
	}
	utils.RenderTemplate(w, "dashboard.html", pageVariables)
}
