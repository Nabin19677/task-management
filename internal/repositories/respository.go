package repositories

import (
	"database/sql"
	"sync"
)

var (
	once sync.Once
	r    *Repository
)

type Repository struct {
	User *UserRepository
	Task *TaskRepository
}

func InitRepositories(db *sql.DB) *Repository {
	once.Do(func() {
		r = &Repository{
			User: NewUserRepository(db),
			Task: NewTaskRepository(db),
		}
	})
	return r
}

func GetRepository() *Repository {
	return r
}
