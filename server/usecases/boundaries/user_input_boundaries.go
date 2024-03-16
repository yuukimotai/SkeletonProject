package boundaries

import (
	"obserbooks/usecases/requests"
	"obserbooks/usecases/responses"
)

type UserInputBoundary interface {
	CreateUser(*requests.UserCreateRequest) (*responses.UserCreateResponse, error)
	GetUser(*requests.UserGetRequest) (*responses.UserGetResponse, error)
}