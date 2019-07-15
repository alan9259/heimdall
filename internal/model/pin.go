package model

import "time"

type Pin struct {
	ID        int    `gorm:"AUTO_INCREMENT;not null;primary_key" json:"id,omitempty"`
	AccountID int    `gorm:"not null;primary_key" json:"account_id,omitempty"`
	Pin       int    `gorm:"not null" json:"pin,omitempty"`
	Purpose   string `json:"purpose,omitempty"`

	VerifiedAt time.Time
	ExpiredAt  time.Time
	CreatedAt  time.Time
}
