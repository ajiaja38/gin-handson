package controller

import (
	"net/http"
	"res-gin/src/dto"
	"res-gin/src/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var userDto dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	user, err := c.userService.CreateUser(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"status":  true,
		"message": "User created successfully",
		"data":    user,
	})
}

func (c *UserController) GetAllUsersHandler(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"status":  false,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  true,
		"message": "Get all users successfully",
		"data":    users,
	})
}

func (c *UserController) GetUserByIdHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	user, err := c.userService.GetUserById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  true,
		"message": "Get user by id successfully",
		"data":    user,
	})
}
