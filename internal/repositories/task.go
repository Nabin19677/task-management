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

func (tr *TaskRepository) GetTableName() string {
	return "tasks"
}

func (tr *TaskRepository) GetTask(taskID int) (*models.Task, error) {
	query := "SELECT id, title, description, status FROM " + tr.GetTableName() + " WHERE id = $1 ;"
	row := tr.db.QueryRow(query, taskID)

	task := &models.Task{}
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Insert(newTask models.NewTask) (bool, error) {
	query := "INSERT INTO " + tr.GetTableName() + " (title, description, due_date, status, manager_id, assignee_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := tr.db.Exec(query, newTask.Title, newTask.Description, newTask.DueDate, newTask.Status, newTask.ManagerID, newTask.AssigneeID)

	if err != nil {
		log.Println("insert failed:", err)
		return false, err
	}

	return true, nil

}

func (tr *TaskRepository) Update(updatedTask models.Task) (bool, error) {
	query := "UPDATE " + tr.GetTableName() + " SET title = $1, description = $2, due_date = $3, status = $4 WHERE id = $5"
	_, err := tr.db.Exec(query, updatedTask.Title, updatedTask.Description, updatedTask.DueDate, updatedTask.Status, updatedTask.ID)

	if err != nil {
		log.Println("update failed:", err)
		return false, err
	}

	return true, nil
}

func (tr *TaskRepository) GetTasksByManager(managerId int) ([]*models.Task, error) {
	var tasks []*models.Task

	query := "SELECT id, title, description, status FROM " + tr.GetTableName() + " WHERE manager_id = $1 ;"
	rows, err := tr.db.Query(query, managerId)

	if err != nil {
		log.Fatal(err)

	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (ur *TaskRepository) GetTasksByAssignee(assigneeId int) ([]*models.Task, error) {
	var tasks []*models.Task

	query := "SELECT id, title, description, status FROM " + ur.GetTableName() + " WHERE assignee_id = $1 ;"
	rows, err := ur.db.Query(query, assigneeId)

	if err != nil {
		log.Fatal(err)

	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskRepository) Delete(taskID int) error {
	query := "DELETE FROM " + tr.GetTableName() + " WHERE id = $1"
	_, err := tr.db.Exec(query, taskID)
	if err != nil {
		log.Println("delete failed:", err)
		return err
	}
	return nil
}
