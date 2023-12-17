package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Status      string
	AssigneeID  int
	ManagerID   int
}

type NewTask struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	DueDate     string `json:"dueData" db:"due_data"`
	Status      string `json:"status" db:"status"`
	AssigneeID  int    `json:"assigneeId" db:"assignee_id"`
	ManagerID   int    `json:"managerId" db:"manager_id"`
}
