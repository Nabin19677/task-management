package handlers

import (
	"net/http"

	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func SuccessHandler(w http.ResponseWriter, message string) {
	pageVariables := types.PageVariables{
		Title: "Success",
		Data: map[string]interface{}{
			"Message": message,
		},
	}
	utils.RenderTemplate(w, "success.html", pageVariables)
}

func FailHandler(w http.ResponseWriter, message string) {
	pageVariables := types.PageVariables{
		Title: "Failed",
		Data: map[string]interface{}{
			"Message": message,
		},
	}
	utils.RenderTemplate(w, "fail.html", pageVariables)
}
