package repository

import (
	"database/sql"
	"fmt"

	"github.com/dyd40000/myapi/models"
)

const (
	articleNumPerPage = 5
)

// 新規記事とデータベースにinsertする関数
// ->データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// クエリの定義
	const sqlStr = `
		insert into article(title,contents,username,nice,created_at) values(?,?,?,0,now());
	`
	// 戻り値の定義
	newArticle := models.Article{
		Title:    article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
	}
	// db.Execでisnert実行
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	// LastInsertIDを取得して
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	return newArticle, nil
}

// 変数pageで指定されたページに表示する記事一覧をデータベースから取得する関数
// ->データベースから取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	// select文宣言
	const sqlStr = `
		select * from article 
		limit ? offset ?;
		`
	// offsetにはどのページまでSKIPするかを指定する。
	// 3ページが指定されたら、2ページまで飛ばせば良い。つまり、((指定されたページ-1)*1ページごとの記事数)がoffsetに入ればよい
	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return []models.Article{}, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.UserName)
		articleArray = append(articleArray, article)
	}

	// 指定したpageのArticleをスライス型で取得する
	return articleArray, nil
}

// 記事IDを指定して、記事データを取得する関数
// ->取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	// select文宣言
	const selectArticleStr = `
		select *
		from article
		where article_id = ?;
	`

	// QueryRow > Scan(&p)
	row := db.QueryRow(selectArticleStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	var createdTime sql.NullTime
	if err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime); err != nil {
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	return article, nil
}

// いいねの数をupdateする関数
// ->発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	// 1.begin
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// 2.select nicenum
	const selectNiceNumStr = `
		select nice 
		from article
		where article_id = ?;
	`
	row := tx.QueryRow(selectNiceNumStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 3.insert nicenum+1
	const updateNiceNumStr = `
		update article 
		set nice = ?
		where article_id = ?;
	`
	_, err = tx.Exec(updateNiceNumStr, niceNum+1, articleID)
	if err != nil {
		fmt.Println("updateでおちた")
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
