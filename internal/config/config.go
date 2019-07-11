package config

import (
	"miu-auth-api-v1/internal/model"
)

type Store interface {
	GetApiKey(string) (*model.Config, error)
}
