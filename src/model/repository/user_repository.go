package repository

import (
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *resterror.RestError)
}
