package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           int32     `gorm:"AUTO_INCREMENT;not null" json:"id,omitempty"`
	EmailAddress string    `gorm:"type:varchar(50);unique_index:idx_account_email;not null" json:"email_address,omitempty"`
	FirstName    string    `gorm:"type:varchar(100);not null" json:"first_name,omitempty"`
	LastName     string    `gorm:"type:varchar(100);" json:"last_name,omitempty"`
	Password     string    `gorm:"type:varchar(100);" json:"password,omitempty"`
	PhoneNumber  string    `gorm:"type:varchar(100);not null" json:"phone_number,omitempty"`
	DateOfBirth  time.Time `json:"date_of_birth,omitempty"`

	Gender   Gender
	GenderID int32

	CreatedAt    time.Time
	VerifiedAt   time.Time
	LastLogin    time.Time
	LastModified time.Time

	Devices []Device `json:"device_list,omitempty"`

	Locations []Location `gorm:"many2many:account_devices;" json:"location_list,omitempty"`
}

func (a *Account) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (a *Account) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
	return err == nil
}
