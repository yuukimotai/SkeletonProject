package controllers

import (
	"fmt"
	"obserbooks/interfaces"
	"obserbooks/usecases/boundaries"
	"obserbooks/usecases/requests"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*interfaces.Controller
	userInputBoundary boundaries.UserInputBoundary
}

// NewUserController コンストラクタ
func NewUserController(userInputBoundary boundaries.UserInputBoundary) *UserController {

	return &UserController{
		Controller:        &interfaces.Controller{},
		userInputBoundary: userInputBoundary,
	}
}

// GetUser 単一のユーザを取得します．
func (uc *UserController) GetUser(context *gin.Context) {
	request := &requests.UserGetRequest{}

	response, err := uc.userInputBoundary.GetUser(request)

	if err != nil {
		uc.ErrorJSON(context, 400, []string{err.Error()})
		return
	}

	context.JSON(200, response)
	return
}

// GetUsers 複数のユーザを取得します．
func (uc *UserController) GetUsers(context *gin.Context) {
}

// CreateUser ユーザを作成します．
func (uc *UserController) CreateUser(context *gin.Context) {
	request := &requests.UserCreateRequest{}

	err := context.ShouldBindJSON(&request)
	fmt.Println(err)

	userCreateResponse, err := uc.userInputBoundary.CreateUser(request)

	if err != nil {
		uc.ErrorJSON(context, 400, []string{err.Error()})
		return
	}

	context.JSON(200, userCreateResponse)
	return
}

// UpdateUser ユーザを更新します．
func (uc *UserController) UpdateUser(context *gin.Context) {
}

// DeleteUser ユーザを削除します．
func (uc *UserController) DeleteUser(context *gin.Context) {
}
