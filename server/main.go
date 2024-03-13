package main

import (
	database "obserbooks/infrastructure/mysql"
	"obserbooks/infrastructure/routers"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	db, err := database.NewDB()

	if err != nil {
		panic(err)
	}

	// 最後にデータベースとの接続を切断します．
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			//log.Log().Error(err.Error())
		}
	}(db)

	router := routers.NewRouter(gin.Default(), db)

	err = router.Run()

	if err != nil {
		//log.Log().Error(err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}
