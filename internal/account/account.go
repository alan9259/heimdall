package account

import (
	"miu-auth-api-v1/internal/model"
)

type Store interface {
	Create(*model.Account) error
	Update(*model.Account) error
	GetByID(uint) (*model.Account, error)
	GetByEmail(string) (*model.Account, error)
}
