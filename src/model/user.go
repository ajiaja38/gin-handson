package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"type:varchar(255); not null; unique" json:"email"`
	Username  string    `gorm:"type:varchar(255); not null" json:"username"`
	Password  string    `gorm:"type:varchar(255); not null" json:"password"`
	Role      string    `gorm:"type:varchar(255); not null" json:"role"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (user *Users) BeforeCreate(db *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
