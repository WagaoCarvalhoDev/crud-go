package service

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
)

func (ud *userDomainService) FindUser(string) (
	model.UserDomainInterface,
	*resterror.RestError,
) {
	return nil, nil
}
