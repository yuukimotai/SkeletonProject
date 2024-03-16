package routers

import (
	database "obserbooks/infrastructure/mysql"
	repositories "obserbooks/infrastructure/repositories/authentication"
	controller "obserbooks/interfaces/controllers"
	"obserbooks/usecases/interactors"

	"github.com/gin-gonic/gin"
)

// UserRouter ユーザに関してルーティングを実行します．
func AuthenticationRouter(router *gin.Engine, db *database.DB) {
	authenticationRouter := router.Group("/authentication")
	{
		c := controller.NewAuthenticationController(interactors.NewAuthenticationInteractor(repositories.NewAuthenticationRepository(db)))
		authenticationRouter.POST("/register", c.Register)
		authenticationRouter.POST("/login", c.Login)
	}
}
