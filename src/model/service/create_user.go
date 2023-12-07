package service

import (
	"crudgo/src/configuration/logger"
	resterror "crudgo/src/configuration/rest_error"
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
