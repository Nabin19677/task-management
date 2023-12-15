package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Status      string
	Assignee    string
}
