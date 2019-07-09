package model

type Device struct {
	ID              int64  `gorm:"AUTO_INCREMENT;not null" json:"device_id,omitempty"`
	SerialNumber    string `gorm:"type:varchar(100);" json:"device_sn,omitempty"`
	Name            string `gorm:"type:varchar(200);" json:"device_name,omitempty"`
	OperatingSystem string `gorm:"type:varchar(100);" json:"device_os,omitempty"`
}
