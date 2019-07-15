package revokedToken

import (
	"miu-auth-api-v1/internal/model"
)

type Store interface {
	Create(*model.RevokedToken) error
	Delete(*model.RevokedToken) error
}
