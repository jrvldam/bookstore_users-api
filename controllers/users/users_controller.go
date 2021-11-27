package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jrvldam/bookstore_users-api/domain/users"
	"github.com/jrvldam/bookstore_users-api/services"
	"github.com/jrvldam/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("User id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.GetUser(id)

	if getErr != nil {
		restErr := errors.NewNotFoundError(fmt.Sprintf("User id %d, not found", id))
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
