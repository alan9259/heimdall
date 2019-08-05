package pin

import "heimdall/internal/model"

type Store interface {
	Create(*model.Pin) error
	GetByCompositeKey(string, int32) (*model.Pin, error)
	RemovePin(string, int32) error
	GetCurrentPins(string) []*model.Pin
}
