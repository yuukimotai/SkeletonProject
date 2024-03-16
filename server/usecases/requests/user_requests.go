package requests

import (
	"github.com/google/uuid"
)

type UserGetRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserCreateRequest struct {
	UserId   uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	UserId uuid.UUID `json:"id" binding:"required,min=1"`
}

type UserDeleteRequest struct {
	UserId uuid.UUID `json:"id" binding:"required,min=1"`
}
