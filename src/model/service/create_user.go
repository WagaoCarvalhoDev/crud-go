package service

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/configuration/rest_error/logger"
	"crudgo/src/model"
	"fmt"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *resterror.RestError {
	logger.Info("Init createUser model ", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
