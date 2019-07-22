package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/google/uuid"

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

func (rts *RevokedTokenStore) GetByJti(jti uuid.UUID) (*model.RevokedToken, error) {
	var m model.RevokedToken
	if err := rts.db.Where(&model.RevokedToken{Jti: jti}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
