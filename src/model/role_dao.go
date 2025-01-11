package model

import (
	"fmt"
	"res-gin/src/enum"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string     `gorm:"type:varchar(255); not null;primary_key" json:"id"`
	Role      enum.ERole `gorm:"type:varchar(255); not null;" json:"role"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
}

func (r *Role) BeforeCreate(db *gorm.DB) (err error) {
	var lastRole Role

	if err := db.Last(&lastRole).Error; err != nil {
		return err
	}

	var lastID int
	_, err = fmt.Sscanf(lastRole.ID, "RO-%d", &lastID)

	if err != nil {
		return err
	}

	r.ID = fmt.Sprintf("RO-%d", lastID+1)

	return nil
}
