package utils

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func ParseTemplate(templateContent string, data interface{}) string {
	tmpl, err := template.New("emailTemplate").Parse(templateContent)
	if err != nil {
		log.Fatal(err)
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		log.Fatal(err)
	}

	return result.String()
}
