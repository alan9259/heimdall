package revokedToken

import (
	"heimdall/internal/model"

	"github.com/google/uuid"
)

type Store interface {
	Create(*model.RevokedToken) error
	Delete(*model.RevokedToken) error
	GetByJti(uuid.UUID) (*model.RevokedToken, error)
}
