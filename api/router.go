package api

import (
	"database/sql"
	"net/http"

	"github.com/dyd40000/myapi/controller"
	"github.com/dyd40000/myapi/service"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	// サーバー全体で使用する構造体を作成
	// 　1.サービス構造体
	// 　2.コントローラー構造体
	s := service.NewMyAppService(db)
	ac := controller.NewArticleController(s)
	cc := controller.NewCommentController(s)

	r := mux.NewRouter()

	r.HandleFunc("/hello", ac.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", ac.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", ac.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", ac.ArticleDetailHandler).Methods(http.MethodGet) // article/IDでURLを登録する
	r.HandleFunc("/article/nice", ac.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cc.PostCommentHandler).Methods(http.MethodPost)

	return r
}
