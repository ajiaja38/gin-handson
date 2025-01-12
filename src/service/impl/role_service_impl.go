package impl

import (
	"fmt"
	"res-gin/src/enum"
	"res-gin/src/model"
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

func (s *RoleServiceImpl) GetOrSaveRole(role enum.ERole) (*model.Role, error) {
	var existRole model.Role

	err := s.db.Where("role = ?", role).First(&existRole).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newRole := &model.Role{
				Role: role,
			}

			fmt.Print(newRole)

			if err := s.db.Create(&newRole).Error; err != nil {
				return nil, err
			}

			return newRole, nil
		}

		return nil, err
	}

	return &existRole, nil
}
