package model

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/configuration/rest_error/logger"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *resterror.RestError {
	logger.Info("Init createUser model ", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud.Password)

	return nil
}
