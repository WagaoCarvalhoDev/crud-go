package model

import "golang.org/x/crypto/bcrypt"

func (u *userDomain) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hash)
	return nil
}
