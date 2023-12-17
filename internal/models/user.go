package models

import (
	"strconv"
	"time"

	"anilkhadka.com.np/task-management/conf"
	"anilkhadka.com.np/task-management/internal/types"
	"github.com/dgrijalva/jwt-go"
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

func (u *User) GenToken() (*AuthToken, error) {
	expireAt := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expireAt.Unix(),
		Id:        strconv.Itoa(u.UserID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    conf.EnvConfigs.JwtIssuer,
	})

	accessToken, err := token.SignedString([]byte(conf.EnvConfigs.JwtSecret))

	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpireAt:    expireAt.UTC().String(),
	}, nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)

	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
