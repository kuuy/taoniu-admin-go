package repositories

import (
	"errors"

	"gorm.io/gorm"

	"taoniu.admin.local/permissions/models"
)

type RolesRepository struct {
	Db *gorm.DB
}

func (r *RolesRepository) Get(email string) *models.Role {
	var entity models.Role
	result := r.Db.Where(
		"email=?",
		email,
	).Take(&entity)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &entity
}
