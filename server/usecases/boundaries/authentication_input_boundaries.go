package boundaries

import (
	"obserbooks/usecases/requests"
	"obserbooks/usecases/responses"

	"github.com/gin-gonic/gin"
)

type AuthenticationInputBoundary interface {
	Register(*requests.RegisterRequest, *gin.Context) (*responses.RegisterResponse, error)
	Login(*requests.LoginRequest, *gin.Context) *responses.LoginResponse
}
