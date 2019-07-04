package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           uint      `json:"id,omitempty"`
	EmailAddress string    `json:"email_address,omitempty"`
	Password     string    `json:"password,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Location     *Location `json:"location,omitempty"`
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
