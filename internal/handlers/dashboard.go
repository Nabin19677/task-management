package handlers

import (
	"net/http"

	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	pageVariables := types.PageVariables{
		Title: "Dashboard",
		Data:  nil,
	}
	utils.RenderTemplate(w, "dashboard.html", pageVariables)
}
