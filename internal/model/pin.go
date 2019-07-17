package model

import "time"

type Pin struct {
	EmailAddress string `gorm:"type:varchar(50);unique_index:idx_account_email;not null;primary_key" json:"email_address,omitempty"`
	Pin          int32  `gorm:"varchar(4);not null;primary_key" json:"pin,omitempty"`
	Purpose      string `gorm:"varchar(15);not null" json:"purpose,omitempty"`
	ExpiredAt    time.Time
}
