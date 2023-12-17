// repository/repository.go
package repositories

import (
	"database/sql"
	"log"

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

func (ur *UserRepository) GetTableName() string {
	return "users"
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

func (ur *UserRepository) FindByID(userID int) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM " + ur.GetTableName() + " WHERE user_id = $1"
	err := ur.db.QueryRow(query, userID).
		Scan(&user.UserID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM " + ur.GetTableName() + " WHERE email = $1"
	err := ur.db.QueryRow(query, email).
		Scan(&user.UserID, &user.Name, &user.Email, &user.PhoneNumber, &user.Role, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Insert(newUser models.NewUser) (bool, error) {
	newUser.HashPassword(newUser.Password)

	// Use SQL query with placeholders to prevent SQL injection
	query := "INSERT INTO " + ur.GetTableName() + " (name, email, phone_number, password, role) VALUES ($1, $2, $3, $4, $5)"
	_, err := ur.db.Exec(query, newUser.Name, newUser.Email, newUser.PhoneNumber, newUser.Password, newUser.Role)

	if err != nil {
		log.Println("insert failed:", err)
		return false, err
	}

	return true, nil

}
