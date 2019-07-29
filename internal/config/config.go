package config

import (
	"heimdall/internal/model"
)

type Store interface {
	GetApiKey(string) (*model.Config, error)
}
