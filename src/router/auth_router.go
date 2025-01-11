package router

import (
	"res-gin/src/controller"
	"res-gin/src/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(api *gin.RouterGroup, db *gorm.DB) {
	authService := impl.NewAuthServiceImpl(db)
	authController := controller.NewAuthController(authService)

	authRouter := api.Group("auth")
	{
		authRouter.POST("login", authController.LoginHandler)
		authRouter.PATCH("refresh-token", authController.RefreshTokenHandler)
	}
}
