package handlers

import (
	"log"
	"net/http"
	"strconv"

	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Context().Value("user_id").(string))
	if err != nil {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	taskService := services.NewTaskService()
	userService := services.NewUserService()

	user, err := userService.GetByID(userID)

	if err != nil {
		log.Fatal(err)
		return
	}

	if user.Role == types.Manager {
		tasks, _ := taskService.GetTasksByManager(userID)

		pageVariables := types.PageVariables{
			Title: "Manager Dashboard",
			Data: map[string]interface{}{
				"Tasks": tasks,
			},
		}
		utils.RenderTemplate(w, "manager_dashboard.html", pageVariables)
	} else if user.Role == types.Assignee {
		tasks, _ := taskService.GetTasksByAssignee(userID)

		pageVariables := types.PageVariables{
			Title: "Your Tasks",
			Data: map[string]interface{}{
				"Tasks": tasks,
			},
		}
		utils.RenderTemplate(w, "assignee_dashboard.html", pageVariables)
	}

}
