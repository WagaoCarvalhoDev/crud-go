package model

import (
	resterror "crudgo/src/configuration/rest_error"

	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(email, password, name string, age int8) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() error {
	// Gera o hash da senha usando bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Atualiza a senha no objeto UserDomain com o hash gerado
	ud.Password = string(hash)
	return nil
}

type UserDomainInterface interface {
	CreateUser() *resterror.RestError
	UpdateUser(string) *resterror.RestError
	FindUser(string) (UserDomain, *resterror.RestError)
	DeleteUser(string) *resterror.RestError
}
