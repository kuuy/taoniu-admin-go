package account

import (
  "errors"

  "github.com/rs/xid"
  "gorm.io/gorm"

  "taoniu.admin.local/account/common"
  "taoniu.admin.local/account/models"
)

type AdminsRepository struct {
  Db *gorm.DB
}

func (r *AdminsRepository) Get(email string) *models.Admin {
  var entity models.Admin
  result := r.Db.Where(
    "email=?",
    email,
  ).Take(&entity)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return nil
  }

  return &entity
}

func (r *AdminsRepository) Create(email string, password string) error {
  var entity models.Admin
  result := r.Db.Where(
    "email=?",
    email,
  ).Take(&entity)
  if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return errors.New("user already exists")
  }
  salt := common.GenerateSalt(16)
  hashedPassword := common.GeneratePassword(password, salt)

  entity = models.Admin{
    ID:       xid.New().String(),
    Email:    email,
    Password: hashedPassword,
    Salt:     salt,
    Status:   1,
  }
  r.Db.Create(&entity)

  return nil
}
