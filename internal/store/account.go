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

func (us *AccountStore) GetByEmail(e string) (*model.Account, error) {
	// var m model.Account
	// if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
	// 	if gorm.IsRecordNotFoundError(err) {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }
	// return &m, nil

	a := model.Account{
		ID:           123,
		EmailAddress: "alan9259@gmail.com",
		Password:     "12345",
		FirstName:    "alan",
		LastName:     "wang",
	}

	if err := us.db.Where(a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}
