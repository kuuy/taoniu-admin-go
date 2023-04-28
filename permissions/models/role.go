package models

import (
	"time"
)

type Role struct {
	ID        string    `gorm:"size:20;primaryKey"`
	Name      string    `gorm:"size:64;not null;uniqueIndex"`
	Status    int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (m *Role) TableName() string {
	return "admin_roles"
}
