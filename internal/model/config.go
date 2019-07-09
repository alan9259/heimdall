package model

type Config struct {
	ID    int32  `gorm:"AUTO_INCREMENT;not null" json:"config_id,omitempty"`
	Name  string `gorm:"type:varchar(100);not null;index:idx_config_name;not null" json:"config_name,omitempty"`
	Type  string `gorm:"type:varchar(100);" json:"config_type,omitempty"`
	Value string `gorm:"type:varchar(5000);not null" json:"config_value,omitempty"`
}
