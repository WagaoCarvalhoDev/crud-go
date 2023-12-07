package repository

import (
	"context"
	"crudgo/src/configuration/logger"
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
	"os"
)

const (
	mongo_db_user = "MONGODB_USERS_DB"
)

func (u *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterror.RestError) {
	logger.Info("Init createUser repository")

	collection_name := os.Getenv(mongo_db_user)

	collection := u.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, resterror.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, resterror.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}
