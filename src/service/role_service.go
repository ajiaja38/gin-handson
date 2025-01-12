package service

import (
	"res-gin/src/enum"
	"res-gin/src/model"
)

type RoleService interface {
	GetOrSaveRole(role enum.ERole) (*model.Role, error)
}
