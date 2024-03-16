package controllers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"obserbooks/interfaces"
	"obserbooks/usecases/boundaries"
	"obserbooks/usecases/requests"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	*interfaces.Controller
	authenticationInputBoundary boundaries.AuthenticationInputBoundary
}

func NewAuthenticationController(authenticationInputBoundary boundaries.AuthenticationInputBoundary) *AuthenticationController {

	return &AuthenticationController{
		Controller:                  &interfaces.Controller{},
		authenticationInputBoundary: authenticationInputBoundary,
	}
}

func (ac AuthenticationController) Register(context *gin.Context) {
	request := &requests.RegisterRequest{}
	err := context.ShouldBindJSON(&request)
	fmt.Println(err)
	response, err := ac.authenticationInputBoundary.Register(request, context)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// エラー処理
		context.JSON(http.StatusInternalServerError, gin.H{"error": "JSON生成エラー"})
	}

	// JSONレスポンスを送信
	if response.StatusCode == 200 {
		context.Data(http.StatusOK, "application/json; charset=utf-8", jsonResponse)
	} else {
		context.Data(http.StatusBadRequest, "application/json; charset=utf-8", jsonResponse)
	}
}

func (ac AuthenticationController) Login(context *gin.Context) {
	request := &requests.LoginRequest{}
	context.ShouldBindJSON(&request)
	response := ac.authenticationInputBoundary.Login(request, context)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// エラー処理
		context.JSON(http.StatusInternalServerError, gin.H{"error": "JSON生成エラー"})
	}

	if response.StatusCode == 200 {
		context.Data(http.StatusOK, "application/json; charset=utf-8", jsonResponse)
	} else {
		context.Data(http.StatusBadRequest, "application/json; charset=utf-8", jsonResponse)
	}
}

func (ac AuthenticationController) Logout(context *gin.Context) {

}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
