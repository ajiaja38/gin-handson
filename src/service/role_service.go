package service

import "res-gin/src/enum"

type RoleService interface {
	GetOrSaveRole(role enum.ERole) (*enum.ERole, error)
}
