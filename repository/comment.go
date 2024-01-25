package repository

import (
	"database/sql"

	"github.com/dyd40000/myapi/models"
)

// 新規コメントをデータベースにinsertする関数
// ->データベースに保存したコメントと、発生したエラーを返り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// const sqlStr
	// db.Exec
	newComment := models.Comment{}
	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
// ->取得したコメントと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	commentList := []models.Comment{}
	return commentList, nil
}
