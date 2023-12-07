package service

import (
	"crudgo/src/configuration/logger"
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"

	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface, *resterror.RestError,
) {
	logger.Info("Init createUser model ", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed succefully",
		zap.String("userId", userDomainRepository.GetId()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
