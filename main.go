package main

import (
	"fmt"

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

	fmt.Println("Que onda amigos")

	router.Run(":8000")
}
