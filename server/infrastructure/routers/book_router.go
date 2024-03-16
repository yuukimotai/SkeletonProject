package routers

import (
	database "obserbooks/infrastructure/mysql"
	repositories "obserbooks/infrastructure/repositories/book"
	controller "obserbooks/interfaces/controllers"
	graphql "obserbooks/usecases/interactors"

	"os"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, db *database.DB) {
	repository := repositories.NewBookRepository(db)
	interactor := *graphql.NewBookInteractor(repository)
	c := controller.NewBookController(interactor)

	router.POST("/query", c.GraphqlHandler(db))
	router.GET("/", c.PlaygroundHandler())
	defaultPort := "8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}
