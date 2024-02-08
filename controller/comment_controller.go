package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dyd40000/myapi/controller/service"
	"github.com/dyd40000/myapi/models"
)

type CommentController struct {
	service service.CommentServicer
}

func NewCommentController(s service.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		fmt.Println(err)
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to post comment\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
