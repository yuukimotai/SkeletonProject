package repositories

import (
	"obserbooks/domain/model"
	usermodel "obserbooks/domain/user"
	"obserbooks/usecases/responses"

	"github.com/gin-gonic/gin"
)

type AuthenticationRepository interface {
	Create(*usermodel.User, *gin.Context) *responses.RegisterResponse
	Login(usermodel.User) *usermodel.User
	FindUser(*usermodel.User, *gin.Context) (*model.User, error)
}
