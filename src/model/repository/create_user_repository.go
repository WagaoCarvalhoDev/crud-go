package repository

import (
	"context"
	"crudgo/src/configuration/logger"
	resterror "crudgo/src/configuration/rest_error"
	"crudgo/src/model"
	"crudgo/src/model/repository/entity/converter"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	mongo_db_user = "MONGODB_USERS_DB"
)

func (u *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterror.RestError) {
	logger.Info("Init createUser repository",
		zap.String("journey", "CreateUser"))

	collection_name := os.Getenv(mongo_db_user)

	collection := u.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to createUser",
			err,
			zap.String("journey", "createUser"))
		return nil, resterror.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
