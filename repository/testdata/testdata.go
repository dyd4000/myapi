package testdata

import "github.com/dyd40000/myapi/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "daniel",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2ndPost",
		Contents: "Second blog post",
		UserName: "daniel",
		NiceNum:  4,
	},
}
