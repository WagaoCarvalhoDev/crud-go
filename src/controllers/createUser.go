package controllers

import (
	resterrors "crudgo/src/configurations/rest_errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := resterrors.NewBadRequestError("Chamada de rota errada")
	c.JSON(err.Code, err)
}
