package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

func (as *AccountStore) Create(a *model.Account) error {
	return as.db.Create(a).Error
}

func (as *AccountStore) Update(a *model.Account) error {
	return as.db.Model(a).Update(a).Error
}

func (as *AccountStore) GetByID(id uint) (*model.Account, error) {
	var m model.Account
	if err := as.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccountStore) GetByEmail(e string) (*model.Account, error) {
	var m model.Account
	if err := as.db.Where(&model.Account{EmailAddress: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
