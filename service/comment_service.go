package service

import (
	"github.com/dyd40000/myapi/models"
	"github.com/dyd40000/myapi/repository"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newCommnet, err := repository.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newCommnet, nil
}
