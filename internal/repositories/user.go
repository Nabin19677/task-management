// repository/repository.go
package repositories

import (
	"database/sql"

	"anilkhadka.com.np/task-management/internal/models"
)

// UserRepository handles user-related operations.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository with a database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUser retrieves a user from the database by ID.
func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := "SELECT user_id, name FROM users WHERE user_id = $1 "
	row := ur.db.QueryRow(query, userID)

	user := &models.User{}
	err := row.Scan(&user.UserID, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
