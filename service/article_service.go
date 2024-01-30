package service

import (
	"fmt"

	"github.com/dyd40000/myapi/models"

	"github.com/dyd40000/myapi/repository"
)

/*
サービス層を実装する上での考え方
- まず、どういった関数が必要か考える
- 関数のシグネチャ（名前・引数・戻り値）から考える
- 引数は、レポジトリ層で操作したいキー情報
- 戻り値はハンドラ層に返したい情報
*/
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repository.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repository.SelectCommentList(db, articleID)
	fmt.Println("for-rangeで取得したcommentsの中身をみる")
	for _, comment := range comments {
		fmt.Printf("comment.message:%s\n", comment.Message)
	}
	if err != nil {
		return models.Article{}, err
	}

	article.Comments = append(article.Comments, comments...)

	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, nil
	}
	defer db.Close()

	newArticle, err := repository.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repository.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, nil
	}
	defer db.Close()

	if err := repository.UpdateNiceNum(db, article.ID); err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum,
		CreatedAt: article.CreatedAt,
	}, nil

}
