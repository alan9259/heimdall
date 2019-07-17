package store

import (
	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
)

type PinStore struct {
	db *gorm.DB
}

func NewPinStore(db *gorm.DB) *PinStore {
	return &PinStore{
		db: db,
	}
}

func (ps *PinStore) Create(p *model.Pin) error {
	return ps.db.Create(p).Error
}

func (ps *PinStore) GetByCompositeKey(e string, p int32) (*model.Pin, error) {
	var pn model.Pin
	if err := ps.db.Where(&model.Pin{EmailAddress: e, Pin: p}).First(&pn).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &pn, nil
}
