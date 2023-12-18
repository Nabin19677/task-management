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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService()

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		email := r.FormValue("username")
		password := r.FormValue("password")

		login := &models.LoginInput{
			Email:    email,
			Password: password,
		}

		auth, err := userService.LoginUser(login)
		if err != nil {
			log.Println(err)
		}

		if auth != nil {
			// Set the token as a cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    auth.AuthToken.AccessToken,
				HttpOnly: true,
			})

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		} else {
			FailHandler(w, "Login Failed")
		}

	} else if r.Method == http.MethodGet {
		HomeHandler(w, r)
	}

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService()
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
			log.Println(err)
		}

		if success {
			SuccessHandler(w, "Signup Successful")
		} else {
			FailHandler(w, "Signup Failed")
		}

	} else if r.Method == http.MethodGet {
		pageVariables := types.PageVariables{
			Title: "Signup",
			Data:  nil,
		}
		utils.RenderTemplate(w, "signup.html", pageVariables)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "auth_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
