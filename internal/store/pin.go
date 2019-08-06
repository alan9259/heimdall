package store

import (
	"heimdall/internal/model"

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

func (ps *PinStore) RemovePin(e string, p int32) error {
	var pn model.Pin
	if err := ps.db.Where(&model.Pin{EmailAddress: e, Pin: p}).Delete(&pn).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PinStore) GetCurrentPins(e string) (*[]model.Pin, error) {
	var pn []model.Pin
	if err := ps.db.Where(&model.Pin{EmailAddress: e}).Find(&pn).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &pn, nil
}
