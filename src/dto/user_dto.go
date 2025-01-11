package dto

import "res-gin/src/enum"

type CreateUserDTO struct {
	Email    string     `json:"email" validate:"required,email"`
	Username string     `json:"username" validate:"required,min=3"`
	Password string     `json:"password" validate:"required,min=6"`
	Role     enum.ERole `json:"role" validate:"required,min=3"`
}
