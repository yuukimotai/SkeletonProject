package routers

import (
	database "obserbooks/infrastructure/mysql"
	repositories "obserbooks/infrastructure/repositories/user"
	controller "obserbooks/interfaces/controllers"
	"obserbooks/usecases/interactors"

	"github.com/gin-gonic/gin"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *database.DB) {
	userRouter := router.Group("/users")
	{
		c := controller.NewUserController(interactors.NewUserInteractor(repositories.NewUserRepository(db)))
		userRouter.GET("/:email", c.GetUser)
		userRouter.POST("", c.CreateUser)
	}
}
