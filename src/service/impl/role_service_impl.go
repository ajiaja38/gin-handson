package impl

import (
	"res-gin/src/enum"
	"res-gin/src/service"

	"gorm.io/gorm"
)

type RoleServiceImpl struct {
	db *gorm.DB
}

func NewRoleServiceImpl(db *gorm.DB) service.RoleService {
	return &RoleServiceImpl{
		db: db,
	}
}

func (r *RoleServiceImpl) GetOrSaveRole(role enum.ERole) (*enum.ERole, error) {
	panic("unimplemented")
}
