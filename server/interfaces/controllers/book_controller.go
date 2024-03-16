package controllers

import (
	"net/http"
	"obserbooks/domain/graph"
	database "obserbooks/infrastructure/mysql"
	"obserbooks/interfaces"
	graphql "obserbooks/interfaces/graph"
	"obserbooks/usecases/interactors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	*interfaces.Controller
	db       *database.DB
	resolver graphql.Resolver
}

func NewBookController(bookinteractor interactors.BookInteractor) *BookController {
	return &BookController{
		Controller: &interfaces.Controller{},
	}
}

func (bc *BookController) GraphqlHandler(db *database.DB) gin.HandlerFunc {
	h := graphqlHandler(db)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (bc *BookController) PlaygroundHandler() gin.HandlerFunc {
	h := playgroundHandler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(db *database.DB) *handler.Server {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graphql.Resolver{DB: db}}))

	return h
}

// Defining the Playground handler
func playgroundHandler() http.Handler {
	h := playground.Handler("GraphQL", "/query")

	return h
}
