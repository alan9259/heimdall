package pin

import "heimdall/internal/model"

type Store interface {
	Create(*model.Pin) error
	GetByCompositeKey(string, int32) (*model.Pin, error)
}
