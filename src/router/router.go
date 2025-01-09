package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(api *gin.RouterGroup, db *gorm.DB) {
	RegisterUserRoutes(api, db)
}
