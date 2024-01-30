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
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repository.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repository.SelectCommentList(s.db, articleID)
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

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repository.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repository.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	if err := repository.UpdateNiceNum(s.db, article.ID); err != nil {
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
