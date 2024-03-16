package responses

import (
	"github.com/google/uuid"
)

type UserGetResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateResponse struct {
	UserId uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
}

type UserUpdateResponse struct {
	UserId uuid.UUID `json:"id"`
	Name   string    `json:"name"`
}

type UserDeleteResponse struct {
	UserId uuid.UUID `json:"id"`
}
