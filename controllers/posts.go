package controllers

import (
	"fmt"

	"github.com/darcops/transcarent/tools"
)

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getPostsByUserId(userId int64) ([]Post, error) {
	var posts []Post

	url := fmt.Sprintf(basePath+"/posts?userId=%v", userId)

	if err := tools.DoGetRequest(url, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
