package controllers

const (
	basePath = "https://jsonplaceholder.typicode.com"
)

type UserPosts struct {
	UserId    int64 `json:"id"`
	*UserInfo `json:"user_info"`
	Posts     []Post `json:"posts"`
}

func (u *UserPosts) Get() error {
	u.UserInfo = &UserInfo{
		Id: u.UserId,
	}
	if err := u.UserInfo.Get(); err != nil {
		return err
	}

	posts, err := getPostsByUserId(u.Id)
	if err != nil {
		return err
	}

	u.Posts = posts
	return nil
}
