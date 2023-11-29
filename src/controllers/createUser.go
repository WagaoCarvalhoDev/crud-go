package controllers

import (
	resterrors "crudgo/src/configurations/rest_errors"
	"crudgo/src/controllers/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterrors.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
