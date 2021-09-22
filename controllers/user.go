package controllers

import (
	"fmt"

	"github.com/darcops/transcarent/tools"
)

type UserInfo struct {
	Id       int64  `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *UserInfo) Get() error {
	url := fmt.Sprintf(basePath+"/users/%v", u.Id)

	if err := tools.DoGetRequest(url, u); err != nil {
		return err
	}

	return nil
}
