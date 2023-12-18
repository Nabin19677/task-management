package repositories

import (
	"database/sql"
	"log"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetTableName() string {
	return "users"
}

func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := "SELECT user_id, name, role FROM users WHERE user_id = $1 "
	row := ur.db.QueryRow(query, userID)

	user := &models.User{}
	err := row.Scan(&user.UserID, &user.Name, &user.Role)
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

	query := "INSERT INTO " + ur.GetTableName() + " (name, email, phone_number, password, role) VALUES ($1, $2, $3, $4, $5)"
	_, err := ur.db.Exec(query, newUser.Name, newUser.Email, newUser.PhoneNumber, newUser.Password, newUser.Role)

	if err != nil {
		log.Println("insert failed:", err)
		return false, err
	}

	return true, nil

}

func (ur *UserRepository) Find() ([]*models.PublicUser, error) {
	var users []*models.PublicUser

	query := "SELECT user_id, name, email, phone_number, role FROM " + ur.GetTableName()

	rows, err := ur.db.Query(query)
	if err != nil {
		log.Println("find failed:", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.PublicUser
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.PhoneNumber, &user.Role)
		if err != nil {
			log.Println("scan failed:", err)
			return users, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error:", err)
		return users, err
	}

	return users, nil
}

func (ur *UserRepository) FindByRole(role types.Role) ([]*models.PublicUser, error) {
	var users []*models.PublicUser

	query := "SELECT user_id, name, email, phone_number, role FROM " + ur.GetTableName() + " WHERE role = $1 ;"

	rows, err := ur.db.Query(query, role)
	if err != nil {
		log.Println("find failed:", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.PublicUser
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.PhoneNumber, &user.Role)
		if err != nil {
			log.Println("scan failed:", err)
			return users, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error:", err)
		return users, err
	}

	return users, nil
}
