package repository

import (
	"database/sql"
	"dbsample/models"
)

const (
	articleNumPerPage = 5
)

// 新規記事とデータベースにinsertする関数
// ->データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// クエリの定義
	const sqlStr = `
		isnert into article(title,contents,user_name,nice,created_at) vlaues(?,?,?,0,now());
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
	// offsetには、どのページまで飛ばすか、を指定する。
	// 3ページが指定されたら、2ページまで飛ばせば良い。つまり、((指定されたページ-1)*1ページごとの記事数)
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
func SlectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
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
	if err := row.Scan(&article.Title, &article.Contents, &article.UserName, &createdTime); err != nil {
		return models.Article{}, nil
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	return article, nil
}

// いいねの数をupdateする関数
// ->発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB) error {
	// const sqlStr
	// db.Exec
	return nil
}
