package model

import "time"

type Location struct {
	ID         int64 `gorm:"AUTO_INCREMENT;not null;"`
	AccountID  int32
	City       string `gorm:"type:varchar(20);not null;" json:"city,omitempty"`
	Region     string `gorm:"type:varchar(20);not null;" json:"state,omitempty"`
	Country    string `gorm:"type:varchar(20);not null;" json:"country,omitempty"`
	IPAddress  string `gorm:"type:varchar(20);not null;" json:"ip_address,omitempty"`
	RecordedAt time.Time
}
