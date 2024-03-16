package usermodel

import (
	"github.com/google/uuid"
)

type User struct {
	UserNumber int `gorm:"primaryKey;autoIncrement"`
	UserId     uuid.UUID
	Name       string
	Email      string
	Password   string
	// role     string
}

type Users []*User

func NewUser(u User) *User {
	return &User{
		UserId:   u.UserId,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
