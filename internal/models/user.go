package models

import (
	"anilkhadka.com.np/task-management/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID      int        `json:"userId" db:"user_id"`
	Name        string     `json:"name" db:"name"`
	Email       *string    `json:"email" db:"email" `
	PhoneNumber *string    `json:"phoneNumber" db:"phone_number"`
	Password    string     `json:"password" db:"password"`
	Role        types.Role `json:"role" db:"role"`
}

type NewUser struct {
	Name        string     `json:"name" db:"name" validate:"required"`
	Email       string     `json:"email" db:"email" validate:"required,email"`
	PhoneNumber string     `json:"phoneNumber" db:"phone_number" validate:"required"`
	Password    string     `json:"password" db:"password" validate:"required"`
	Role        types.Role `json:"role" db:"role"`
}

func (u *NewUser) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return err
}
