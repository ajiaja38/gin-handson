package service

import (
	"res-gin/src/dto"
	"res-gin/src/model"
)

type UserService interface {
	CreateUser(userDto *dto.CreateUserDTO) (*model.Users, error)
	GetUserById(id string) (*model.Users, error)
	GetAllUsers() ([]model.Users, error)
	VerifyPassword(hashedPassword, password string) error
}
