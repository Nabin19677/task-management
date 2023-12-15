// repository/repository.go
package repositories

import (
	"database/sql"

	"anilkhadka.com.np/task-management/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (ur *TaskRepository) GetTask(taskID int) (*models.Task, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := ur.db.QueryRow(query, taskID)

	task := &models.Task{}
	err := row.Scan(&task.Title)
	if err != nil {
		return nil, err
	}

	return task, nil
}
