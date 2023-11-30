package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword() error
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.password = string(hash)
	return nil
}
