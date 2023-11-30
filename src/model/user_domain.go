package model

import resterror "crudgo/src/configuration/rest_error"

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *resterror.RestError
	UpdateUser(string, UserDomain) *resterror.RestError
	FindUser(string) (UserDomain, *resterror.RestError)
	DeleteUser(string) *resterror.RestError
}
