package models

import (
	"gorm.io/gorm"
	"taoniu.admin.local/account/models/admin"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

func (m *Admin) AutoMigrate(db *gorm.DB) error {
	db.AutoMigrate(
		&admin.User{},
	)
	return nil
}
