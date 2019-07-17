package pin

import "miu-auth-api-v1/internal/model"

type Store interface {
	Create(*model.Pin) error
	GetByCompositeKey(string, int32) (*model.Pin, error)
}
