package account

import (
	"miu-auth-api-v1/internal/model"
)

type Store interface {
	GetByEmail(string) (*model.Account, error)
	Create(*model.Account) error
	UpdateAccountDetails(*model.Account) error
}
