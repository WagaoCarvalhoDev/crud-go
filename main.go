package main

import (
	"context"
	"crudgo/src/configuration/database/mongodb"
	"crudgo/src/configuration/logger"
	"crudgo/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error(),
		)
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
