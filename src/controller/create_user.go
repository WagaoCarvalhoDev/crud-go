package controller

import (
	"crudgo/src/configuration/rest_error/validation"
	"crudgo/src/controller/model/request"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
