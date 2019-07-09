package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
)

type LocationStore struct {
	db *gorm.DB
}

func NewLocationStore(db *gorm.DB) *LocationStore {
	return &LocationStore{
		db: db,
	}
}

func (ls *LocationStore) Create(l *model.Location) (err error) {
	return ls.db.Create(l).Error
}
