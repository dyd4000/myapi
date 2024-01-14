package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	/* 仕様
	pageの変数が数字だった場合は記事一覧ページのxページ目に表示されるデータを返す
	pageに対応する複数の数字が含まれる場合は最初の値を使用する
	xが数字でなかった場合には400を返す
	クエリパラメーがURLについていなかった場合には、page=1がついているとみなす
	*/

	/* 初手：リクエストに含まれているクエリパラメータを取得するために、1,2を一気にやる
	1.reqの*URLフィールドを取得
	2.Queryメソッドで*URLからクエリパラメータを取得
	*/
	queryMap := req.URL.Query()

	// 何ページ目の記事が欲しいのかを格納するpage変数を定義する
	var page int

	// クエリパラメータに"page"をkeyに持つ値が存在していて、かつクエリパラメータが一つ以上
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータpageに対応する一つ目の値を取得してintに変換
		var err error
		page, err = strconv.Atoi(p[0])

		//数値に変換できない場合＝クエリパラメータが文字列じゃなかった場合はBadRequestで返す
		if err != nil {
			http.Error(w, "Invalid query Parameter", http.StatusBadRequest)
			return
		}
		//パラメータにpage"をkeyに持つ値が存在していない場合
	} else {
		// page=1と同じ処理をする
		page = 1
	}
	resString := fmt.Sprintf("Article List(page%d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment")
}
