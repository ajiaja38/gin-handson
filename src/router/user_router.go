package router

import (
	"res-gin/src/controller"
	"res-gin/src/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(api *gin.RouterGroup, db *gorm.DB) {
	userService := impl.NewUserService(db)
	userController := controller.NewUserController(userService)

	api.POST("/user", userController.CreateUserHandler)
	api.GET("/users", userController.GetAllUsersHandler)
}
