package model

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) GetId() string {
	return u.id
}

func (u *userDomain) SetId(id string) {
	u.id = id
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetAge() int8 {
	return u.age
}
