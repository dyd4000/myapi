package main

import (
	"log"
	"net/http"

	"github.com/dyd40000/myapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// ルータを宣言
	r := mux.NewRouter()

	// パスとハンドラをルータに登録
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet) // article/IDでURLを登録する
	r.HandleFunc("/article/nice", handlers.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")

	// サーバーを起動。ルータを指定
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
