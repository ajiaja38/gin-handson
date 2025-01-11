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
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Roles     []Role    `gorm:"many2many:user_roles; foreignKey:id; joinForeignKey:user_id; joinReferences:role_id;references:ID" json:"roles"`
}

func (user *Users) BeforeCreate(db *gorm.DB) error {
	user.ID = "user-" + uuid.New().String()
	return nil
}
