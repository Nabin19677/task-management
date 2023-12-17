// repository/repository.go
package repositories

import (
	"database/sql"
	"log"

	"anilkhadka.com.np/task-management/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (ur *TaskRepository) GetTableName() string {
	return "tasks"
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

func (ur *TaskRepository) Insert(newTask models.NewTask) (bool, error) {
	query := "INSERT INTO " + ur.GetTableName() + " (title, description, due_date, status, manager_id, assignee_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := ur.db.Exec(query, newTask.Title, newTask.Description, newTask.DueDate, newTask.Status, newTask.ManagerID, newTask.AssigneeID)

	if err != nil {
		log.Println("insert failed:", err)
		return false, err
	}

	return true, nil

}

func (ur *TaskRepository) GetTasksByManager(managerId int) ([]*models.Task, error) {
	var tasks []*models.Task

	query := "SELECT id, title, description FROM " + ur.GetTableName() + " WHERE manager_id = $1 ;"
	rows, err := ur.db.Query(query, managerId)

	if err != nil {
		log.Fatal(err)

	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
