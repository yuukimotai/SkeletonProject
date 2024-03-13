package requests

import "github.com/google/uuid"

type RegisterRequest struct {
	UserId   uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
