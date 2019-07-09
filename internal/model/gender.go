package model

type Gender struct {
	ID   int32  `gorm:"AUTO_INCREMENT;not null" json:"gender_id,omitempty"`
	Name string `gorm:"type:varchar(20);not null;" json:"gender_name,omitempty"`
}
