package routes

import (
	"crudgo/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controllers.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controllers.FindUserByEmail)
	r.POST("/createUser/", controllers.CreateUser)
	r.PUT("/updateUser/:userId", controllers.UpdateUser)
	r.DELETE("/deleteUser/:userId", controllers.DeleteUser)
}
