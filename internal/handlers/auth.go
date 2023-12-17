package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/internal/types"
	"anilkhadka.com.np/task-management/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Your login logic here
	// ...

	// After successful login, redirect to another page
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		HomeHandler(w, r)
	}

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService()
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		phoneNumber := r.FormValue("phoneNumber")
		password := r.FormValue("password")
		role := r.FormValue("role")

		roleNumber, err := strconv.Atoi(role)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		newUser := &models.NewUser{
			Name:        name,
			Email:       email,
			PhoneNumber: phoneNumber,
			Password:    password,
			Role:        roleNumber,
		}

		success, err := userService.CreateUser(newUser)
		if err != nil {
			log.Fatal(err)
		}

		if success {
			SuccessHandler(w, "Signup Successful")
		} else {
			fmt.Println("User creation failed")
		}

	} else if r.Method == http.MethodGet {
		pageVariables := types.PageVariables{
			Title: "Signup",
			Data:  nil,
		}
		utils.RenderTemplate(w, "signup.html", pageVariables)
	}
}
