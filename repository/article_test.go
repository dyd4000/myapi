package repository_test

import (
	"testing"

	"github.com/dyd40000/myapi/models"
	"github.com/dyd40000/myapi/repository"
	"github.com/dyd40000/myapi/repository/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	// 1.テストケース名とテストデータがワンセットになった構造体のスライスの作成
	//   テストデータはつまりテスト結果として期待するデータのこと
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	// for-range文の中でtestを使って期待値と実績値の比較をしていく
	for _, test := range tests {

		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repository.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal()
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID:get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get '%s' but want '%s'\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get '%s' but want '%s'\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}

		})
	}
}

func TestSlectArticleList(t *testing.T) {
	const articleNum = 2

	tests := []struct {
		testTitle string
		pageNum   int
		expected  models.Article
	}{
		{
			testTitle: "Page1Test",
			pageNum:   1,
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "Page2Test",
			pageNum:   2,
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repository.SelectArticleList(testDB, test.pageNum)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != articleNum {
				t.Errorf("ArticleNum: get %d but want %d", len(got), articleNum)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "CleanUpTest3",
		Contents: "This article is almost deleted",
		UserName: "daniel",
	}
	newArticle, err := repository.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != article.Title {
		t.Errorf("new article id expected %s but got %s\n", article.Title, newArticle.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
				delete from article
				where title = ? and contents = ? and username =?;
			`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := testdata.ArticleTestData[0].ID
	expectedNiceNum := testdata.ArticleTestData[0].NiceNum
	err := repository.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repository.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if got.NiceNum != expectedNiceNum+1 {
		t.Errorf("nice num expected %d but got %d", expectedNiceNum+1, got.NiceNum)
	}
}
