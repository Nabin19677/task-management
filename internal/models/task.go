package models

import (
	"anilkhadka.com.np/task-management/internal/types"
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Status      types.TaskStatus
	AssigneeID  int
	ManagerID   int
}

type NewTask struct {
	Title       string           `json:"title" db:"title"`
	Description string           `json:"description" db:"description"`
	DueDate     string           `json:"dueData" db:"due_data"`
	Status      types.TaskStatus `json:"status" db:"status"`
	AssigneeID  int              `json:"assigneeId" db:"assignee_id"`
	ManagerID   int              `json:"managerId" db:"manager_id"`
}
