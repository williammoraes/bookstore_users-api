package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williammoraes/bookstore_users-api/domain/users"
	"github.com/williammoraes/bookstore_users-api/services"
	"github.com/williammoraes/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadResquesError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr)
		return
	} else {
		c.JSON(http.StatusCreated, result)
	}
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
