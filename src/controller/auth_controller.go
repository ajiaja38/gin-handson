package controller

import (
	"net/http"
	"res-gin/src/dto"
	"res-gin/src/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) LoginHandler(ctx *gin.Context) {
	var loginDto dto.LoginDTO

	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	token, err := c.authService.LoginUser(&loginDto)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  true,
		"message": "Login successfully",
		"data":    token,
	})
}

func (c *AuthController) RefreshTokenHandler(ctx *gin.Context) {
	var refreshTokenDto dto.RefreshTokenDTO

	if err := ctx.ShouldBindJSON(&refreshTokenDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	AccessToken, err := c.authService.RefreshToken(&refreshTokenDto)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  true,
		"message": "Refresh token successfully",
		"data":    AccessToken,
	})
}
