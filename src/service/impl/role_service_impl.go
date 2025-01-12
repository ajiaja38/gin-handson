package impl

import (
	"log"
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

	if err := s.db.Where("role = ?", role).First(&existRole).Error; err == nil {
		log.Println("Existing Role Found:", existRole.ID, existRole.Role)
		return &existRole, nil
	}

	newRole := &model.Role{
		Role: role,
	}

	if err := s.db.Create(&newRole).Error; err != nil {
		log.Println("Error Creating Role:", err)
		return nil, err
	}

	log.Println("New Role Created:", newRole.ID, newRole.Role)
	return newRole, nil
}
