package service

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
	"crudgo/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *resterror.RestError)
	UpdateUser(string, model.UserDomainInterface) *resterror.RestError
	FindUser(string) (model.UserDomainInterface, *resterror.RestError)
	DeleteUser(string) *resterror.RestError
}
