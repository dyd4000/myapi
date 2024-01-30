package service

import (
	"github.com/dyd40000/myapi/models"
	"github.com/dyd40000/myapi/repository"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	newCommnet, err := repository.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newCommnet, nil
}
