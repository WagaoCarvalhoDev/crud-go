package service

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *resterror.RestError
	UpdateUser(string, model.UserDomainInterface) *resterror.RestError
	FindUser(string) (model.UserDomainInterface, *resterror.RestError)
	DeleteUser(string) *resterror.RestError
}
