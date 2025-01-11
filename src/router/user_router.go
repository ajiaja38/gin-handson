package router

import (
	"res-gin/src/controller"
	"res-gin/src/middleware"
	"res-gin/src/model"
	"res-gin/src/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(api *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&model.Users{}, &model.Role{})

	userService := impl.NewUserService(db)
	userController := controller.NewUserController(userService)

	userRouter := api.Group("user")
	{
		userRouter.POST("register", userController.CreateUserHandler)
		userRouter.Use(middleware.JWTAuthGuard())
		{
			userRouter.GET("", userController.GetAllUsersHandler)
			userRouter.GET(":id", userController.GetUserByIdHandler)
			userRouter.DELETE(":id", userController.DeleteUserByIdHandler)
		}
	}
}
