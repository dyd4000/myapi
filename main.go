package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dyd40000/myapi/controller"
	"github.com/dyd40000/myapi/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 初期処理
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

	// ルータを宣言
	r := mux.NewRouter()

	s := service.NewMyApiService(db)
	c := controller.NewMyAppController(s)

	// パスとハンドラをルータに登録
	r.HandleFunc("/hello", c.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", c.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", c.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", c.ArticleDetailHandler).Methods(http.MethodGet) // article/IDでURLを登録する
	r.HandleFunc("/article/nice", c.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", c.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")

	// サーバーを起動。ルータを指定
	err = http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
