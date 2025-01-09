package router

import (
	"res-gin/src/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(api *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
	RegisterUserRoutes(api, db)
}
