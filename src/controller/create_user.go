package controller

import (
	"crudgo/src/configuration/rest_error/logger"
	"crudgo/src/configuration/rest_error/validation"
	"crudgo/src/controller/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)
}
