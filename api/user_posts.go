package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darcops/transcarent/controllers"

	"github.com/gin-gonic/gin"
)

const (
	minValidId = 1
	maxValidId = 10
)

func GetUserPosts(c *gin.Context) {
	id := c.Param("id")
	formatedId, _ := strconv.Atoi(id)

	//Stop the process if the given id is invalid.
	//If this project grows, this code could be part of a middleware.
	if formatedId < minValidId || formatedId > maxValidId {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid id. Allowed values are from 1 to 10."))
		return
	}

	userPosts := &controllers.UserPosts{
		UserId: int64(formatedId),
	}

	if err := userPosts.Get(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userPosts)
}
