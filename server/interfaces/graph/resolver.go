package graphql

import (
	"context"
	"fmt"
	"log"
	database "obserbooks/infrastructure/mysql"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Claimsチェックのハンドリングしてない
type Resolver struct {
	DB *database.DB
}

func ValidateJwt(ctx context.Context) (*jwt.Token, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	token := gc.GetHeader("Authorization")
	result, err := validateToken(token)

	return result, err
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	//bytes, err := os.ReadFile("private.pem")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Secret")), nil
	})
	if err != nil {
		return nil, err
	}
	claimsCheckErr := token.Claims.Valid()
	if claimsCheckErr != nil {
		fmt.Println(claimsCheckErr)
		return nil, claimsCheckErr
	}

	return token, nil
}
