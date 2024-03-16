package routers

import (
	"context"
	"obserbooks/infrastructure/middlewares"
	database "obserbooks/infrastructure/mysql"
	graphql "obserbooks/interfaces/graph"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	c        *gin.Context
	router   *gin.Engine
	db       *database.DB
	resolver graphql.Resolver
}

func NewRouter(router *gin.Engine, db *database.DB) *Router {
	return &Router{
		router: router,
		db:     db,
	}
}

func (r *Router) Run() error {
	r.router.Use( //未実装箇所あり
		cors.New(cors.Config{
			// アクセスを許可したいアクセス元
			AllowOrigins: []string{
				"http://localhost:5173",
			},
			// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
			AllowMethods: []string{
				"POST",
				"GET",
			},
			// 許可したいHTTPリクエストヘッダ
			AllowHeaders: []string{
				"Access-Control-Allow-Credentials",
				"Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
			},
			// cookieなどの情報を必要とするかどうか
			// preflightリクエストの結果をキャッシュする時間
			MaxAge: 24 * time.Hour,
		}),
		middlewares.HandleError(),
		GinContextToContextMiddleware())
	UserRouter(r.router, r.db)
	AuthenticationRouter(r.router, r.db)
	BookRouter(r.router, r.db)

	return r.router.Run()
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
