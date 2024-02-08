package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dyd40000/myapi/api"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// サーバー起動時の初期処理として*db.Sqlを作成する
	err := godotenv.Load("./.env")
	if err != nil {
		panic("")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect DB")
	}

	// コントローラーのハンドラメソッドとパスをルータに登録
	r := api.NewRouter(db)

	log.Println("server start at port 8080")

	// サーバーを起動。ルータを指定
	log.Fatal(http.ListenAndServe(":8080", r))
}
