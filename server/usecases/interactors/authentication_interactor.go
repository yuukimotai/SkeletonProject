package interactors

import (
	"fmt"
	"net/http"
	repositories "obserbooks/domain/domain-repositories"
	usermodel "obserbooks/domain/user"
	"obserbooks/usecases/boundaries"
	"obserbooks/usecases/requests"
	"obserbooks/usecases/responses"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TODO 20231018
// Register 新規登録
// FindUser 新規登録時の存在チェック
// Logout jwtの無効化してリダイレクト処理のきっかけ作り
type AuthenticationInteractor struct {
	authenticationRepository repositories.AuthenticationRepository
}

var _ boundaries.AuthenticationInputBoundary = &AuthenticationInteractor{}

func NewAuthenticationInteractor(authenticationRepository repositories.AuthenticationRepository) *AuthenticationInteractor {
	return &AuthenticationInteractor{
		authenticationRepository: authenticationRepository,
	}
}

func (ai *AuthenticationInteractor) Register(registerRequest *requests.RegisterRequest, context *gin.Context) (*responses.RegisterResponse, error) {
	var request usermodel.User

	request.UserId, _ = uuid.NewUUID()
	request.Name = registerRequest.Name
	request.Email = registerRequest.Email

	hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	fmt.Println(string(hashed))
	request.Password = string(hashed)
	user := usermodel.NewUser(request)

	createResult := ai.authenticationRepository.Create(user, context) //ここでステータスコードを返す

	return &responses.RegisterResponse{
		StatusCode: createResult.StatusCode,
		StatusText: createResult.StatusText,
	}, nil
}

func (ai *AuthenticationInteractor) Login(loginRequest *requests.LoginRequest, context *gin.Context) *responses.LoginResponse {
	if loginRequest.Email == "" || loginRequest.Password == "" {
		context.String(http.StatusBadRequest, "bad request")
		return &responses.LoginResponse{}
	}

	var request usermodel.User
	hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	request.Email = loginRequest.Email
	request.Password = string(hashed)
	user := ai.authenticationRepository.Login(request)

	secret := os.Getenv("Secret")
	byteSecret := []byte(secret)
	token, err := generateToken(user, byteSecret)
	if err != nil {
		fmt.Println("Jwt生成に失敗しました")
	}

	response := &responses.LoginResponse{}
	if loginRequest.Email != "" || loginRequest.Password != "" {
		response.StatusCode = 200
		response.StatusText = "ログインしました"
		response.Jwt = token
	}

	return response
}

func generateToken(user *usermodel.User, jwtSecret []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "obserbooksweb", // ドメイン情報入れる
		Subject:   user.UserId.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	token, err := claims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
