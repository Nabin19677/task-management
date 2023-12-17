package types

type TaskStatus string

const (
	StatusTodo       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN PROGRESS"
	StatusInReview   TaskStatus = "IN REVIEW"
	StatusDone       TaskStatus = "DONE"
)
