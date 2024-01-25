package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "first comment1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:        1,
		Title:     "first article",
		Contents:  "This is test article",
		UserName:  "daniel",
		NiceNum:   1,
		Comments:  []Comment{Comment1, Comment2},
		CreatedAt: time.Now(),
	}
	Article2 = Article{
		ID:        1,
		Title:     "first article",
		Contents:  "This is test article",
		UserName:  "daniel",
		NiceNum:   1,
		CreatedAt: time.Now(),
	}
)
