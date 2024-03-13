package repositories

import (
	"fmt"
	repositories "obserbooks/domain/domain-repositories"
	"obserbooks/domain/model"
	usermodel "obserbooks/domain/user"
	"obserbooks/infrastructure/dto"
	database "obserbooks/infrastructure/mysql"
	"obserbooks/usecases/responses"

	"github.com/gin-gonic/gin"
)

type AuthenticationRepository struct {
	db      *database.DB
	context *gin.Context
}

var _ repositories.AuthenticationRepository = &AuthenticationRepository{}

func NewAuthenticationRepository(db *database.DB) repositories.AuthenticationRepository {
	return &AuthenticationRepository{
		db: db,
	}
}

func (ar *AuthenticationRepository) Create(user *usermodel.User, context *gin.Context) *responses.RegisterResponse {
	userDTO := dto.UserDTO{
		UserId:   user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	result := ar.db.FindByEmailAndPassword(userDTO)
	if result.Email != "" || result.Password != "" {
		return &responses.RegisterResponse{
			StatusCode: 400,
			StatusText: "ユーザー情報を登録できません",
		}
	}
	err := ar.db.Create(&userDTO)
	if err != nil {
		return &responses.RegisterResponse{
			StatusCode: 400,
			StatusText: "ユーザー情報を登録できませんでした",
		}
	}

	return &responses.RegisterResponse{
		StatusCode: 200,
		StatusText: "ユーザー情報を登録しました",
	}
}

func (ar *AuthenticationRepository) Login(user usermodel.User) *usermodel.User {
	userDTO := dto.UserDTO{}

	userDTO.Email = user.Email
	userDTO.Password = user.Password

	result, err := ar.db.Find(userDTO)
	fmt.Println(result)
	if err != nil {
		println(err)
	}

	return result
}

func (ar *AuthenticationRepository) FindUser(user *usermodel.User, context *gin.Context) (*model.User, error) {
	userDTO := dto.UserDTO{}

	userDTO.Email = user.Email
	userDTO.Password = user.Password

	result := ar.db.FindByEmailAndPassword(userDTO)

	return result, nil
}
