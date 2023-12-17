package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"

	"anilkhadka.com.np/task-management/internal/types"
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
