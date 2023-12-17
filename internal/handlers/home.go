package handlers

import (
	"net/http"

	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pageVariables := types.PageVariables{
		Title: "Login Page",
	}
	utils.RenderTemplate(w, "index.html", pageVariables)
}
