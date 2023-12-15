package utils

import (
	"net/http"
	"regexp"
)

func RegisterRoute(route string, handler http.HandlerFunc) {
	http.HandleFunc("/"+route, handler)
	http.HandleFunc("/"+route+"/", handler)
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
