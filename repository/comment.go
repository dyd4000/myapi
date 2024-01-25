package repository

import (
	"database/sql"

	"github.com/dyd40000/myapi/models"
)

// 新規コメントをデータベースにinsertする関数
// ->データベースに保存したコメントと、発生したエラーを返り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// const sqlStr
	const sqlStr = `
		isnert into comment() values(?,?,now());
	`
	newCommnet := models.Comment{
		ArticleID: comment.ArticleID,
		Message:   comment.Message,
	}
	// db.Exec
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()
	newCommnet.CommentID = int(id)
	return newCommnet, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
// ->取得したコメントと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	// const
	const sqlStr = `
		select message
		from comment
		where article_id = ?;
	`
	// query
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return []models.Comment{}, err
	}
	defer rows.Close()

	// next>scan
	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}
