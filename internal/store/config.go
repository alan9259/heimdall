package store

import (
	"heimdall/internal/model"

	"github.com/jinzhu/gorm"
)

type ConfigsStore struct {
	db *gorm.DB
}

func NewConfigStore(db *gorm.DB) *ConfigsStore {
	return &ConfigsStore{
		db: db,
	}
}

func (us *ConfigsStore) GetApiKey(e string) (*model.Config, error) {
	var m model.Config
	if err := us.db.Where(&model.Config{Name: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
