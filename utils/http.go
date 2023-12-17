package utils

import (
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"

	"anilkhadka.com.np/task-management/conf"
	"anilkhadka.com.np/task-management/internal/types"
	"github.com/dgrijalva/jwt-go"
)

func RegisterRoute(route string, handler http.HandlerFunc) {
	http.HandleFunc("/"+route, handler)
	http.HandleFunc("/"+route+"/", handler)
}

func RenderTemplate(w http.ResponseWriter, tmplFile string, pageVariables types.PageVariables) {
	tmplPath := filepath.Join("internal/templates", tmplFile)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ExtractIDFromURL(path, prefix string) string {
	re := regexp.MustCompile(prefix + `(\d+)`)
	matches := re.FindStringSubmatch(path)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func ExtractParamsFromURL(path, prefix string) []string {
	re := regexp.MustCompile(prefix + `([^/]+)/(\d+)`)
	matches := re.FindStringSubmatch(path)
	if len(matches) > 2 {
		return matches[1:]
	}
	return nil
}

func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		secret := []byte(conf.EnvConfigs.JwtSecret)
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
