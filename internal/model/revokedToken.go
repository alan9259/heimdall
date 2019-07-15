package model

import (
	"time"
)

type RevokedToken struct {
	AccountId int32  `gorm:"primary_key;"`
	Jti       string `gorm:"primary_key;type:varchar(100);"`
	Token     string `gorm:"type:varchar(100);not null"`
	ExpiredAt time.Time
	RevokedAt time.Time
}
