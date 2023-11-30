package service

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *resterror.RestError {
	return nil
}
