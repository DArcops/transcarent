package controllers

import "sync"

const (
	basePath = "https://jsonplaceholder.typicode.com"
)

type UserPosts struct {
	UserId    int64 `json:"id"`
	*UserInfo `json:"user_info"`
	Posts     []Post `json:"posts"`
}

func (u *UserPosts) Get() error {
	var wg sync.WaitGroup
	resultErr := make(chan error)

	wg.Add(2)

	go func() {
		defer wg.Done()

		u.UserInfo = &UserInfo{
			Id: u.UserId,
		}
		err := u.UserInfo.Get()

		resultErr <- err
	}()

	go func() {
		defer wg.Done()

		posts, err := getPostsByUserId(u.UserId)
		resultErr <- err

		u.Posts = posts
	}()

	go func() {
		wg.Wait()
		close(resultErr)
	}()

	for err := range resultErr {
		if err != nil {
			return err
		}
	}

	return nil
}
