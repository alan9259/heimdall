package model

import (
	"time"

	"github.com/google/uuid"
)

type RevokedToken struct {
	Jti       uuid.UUID `gorm:"primary_key;unique;type:uuid"`
	Token     string    `gorm:"type:varchar;not null"`
	AccountId int32
	ExpiredAt time.Time
	RevokedAt time.Time
}
