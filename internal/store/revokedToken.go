package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
)

type RevokedTokenStore struct {
	db *gorm.DB
}

func NewRevokedTokenStore(db *gorm.DB) *RevokedTokenStore {
	return &RevokedTokenStore{
		db: db,
	}
}

func (rts *RevokedTokenStore) Create(rt *model.RevokedToken) error {
	return rts.db.Create(rt).Error
}

func (rts *RevokedTokenStore) Delete(rt *model.RevokedToken) error {
	return rts.db.Delete(rt).Error
}
