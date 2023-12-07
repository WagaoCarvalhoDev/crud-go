package model

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	SetId(string)

	GetJSONValue() (string, error)

	EncryptPassword() error
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email: email, Password: password, Name: name, Age: age,
	}
}

func (u *userDomain) SetId(id string) {
	u.Id = id
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

type userDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
