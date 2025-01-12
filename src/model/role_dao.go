package model

import (
	"res-gin/src/enum"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        string     `gorm:"type:varchar(255); not null;primary_key" json:"id"`
	Role      enum.ERole `gorm:"type:varchar(255); not null; unique" json:"role"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
}

func (role *Role) BeforeCreate(db *gorm.DB) error {
	role.ID = "RO-" + strings.Split(uuid.New().String(), "-")[1]
	return nil
}
