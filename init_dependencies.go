package main

import (
	"crudgo/src/controller"
	"crudgo/src/model/repository"
	"crudgo/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)

	return controller.NewUserControllerInterface(service)
}
