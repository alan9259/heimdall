package revokedToken

import (
	"miu-auth-api-v1/internal/model"

	"github.com/google/uuid"
)

type Store interface {
	Create(*model.RevokedToken) error
	Delete(*model.RevokedToken) error
	GetByJti(uuid.UUID) (*model.RevokedToken, error)
}
