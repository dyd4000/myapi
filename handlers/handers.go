package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dyd40000/myapi/models"
	"github.com/dyd40000/myapi/service"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// req.Body（ストリーム）を格納する変数
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	newArticle, err := service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "failed to post article\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newArticle)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	/*
		仕様：
		pageの変数が数字だった場合は記事一覧ページのxページ目に表示されるデータを返す
		pageに対応する複数の数字が含まれる場合は最初の値を使用する
		xが数字でなかった場合には400を返す
		クエリパラメーがURLについていなかった場合には、page=1がついているとみなす
	*/

	// reqの*URLフィールドを取得し、
	// Queryメソッドで*URL構造体のRawQueryのクエリパラメータを取得

	queryMap := req.URL.Query()

	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query Parameter\n", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "failed to get article list", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter\n", http.StatusBadRequest)
		return
	}

	article, err := service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to get article\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	article, err := service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "failed to post nice\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		fmt.Println(err)
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	comment, err := service.PostCommentService(reqComment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to post comment\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
