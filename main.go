package main

import (
	"github.com/darcops/transcarent/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiRoutes := router.Group("v1")
	{
		userPosts := apiRoutes.Group("user-posts")
		{
			userPosts.GET(":id", api.GetUserPosts)
		}
	}
	router.Run(":8000")
}
