package database

import (
	"fmt"
	"os"

	"obserbooks/domain/model"
	usermodel "obserbooks/domain/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	connection *gorm.DB
}

func NewDB() (*DB, error) {
	err := godotenv.Load("./infrastructure/mysql/mysql.env")
	if err != nil {
		// .env ファイルが読み込めない場合の処理
	}

	// 環境変数から値を取得
	dsn := fmt.Sprintf(
		"%s:%s@tcp(mysql-db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	conn.AutoMigrate(usermodel.User{}) //後で、テーブルがなかったら作るように書き換え
	conn.AutoMigrate(model.Book{})
	conn.AutoMigrate(model.ReadingMemo{})

	if err != nil {
		return nil, err
	}

	return &DB{
		connection: conn,
	}, nil
}

func ConnectDatabase() {
	dsn := "yuuki:2733ek9se3pe@tcp(mysql-db)/test_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	print(db, err)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
}

func (d *DB) Close() error {
	sqlDb, err := d.connection.DB()

	if err != nil {
		return err
	}

	err = sqlDb.Close()

	if err != nil {
		return err
	}

	return nil
}
