package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func (as *AccountStore) Create(a *model.Account) (err error) {
	return as.db.Create(a).Error
}
