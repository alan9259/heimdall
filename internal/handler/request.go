package handler

import (
	"miu-auth-api-v1/internal/model"
	"time"

	"github.com/labstack/echo"
)

type registerRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

func (r *registerRequest) bind(c echo.Context, a *model.Account) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.EmailAddress = r.EmailAddress
	a.FirstName = r.FirstName
	h, err := a.HashPassword(r.Password)
	if err != nil {
		return err
	}
	a.Password = h
	return nil
}

type loginRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

func (r *loginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type accountUpdateRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	GenderID    int32     `json:"gender_id"`
}

func newAccountUpdateRequest() *accountUpdateRequest {
	return new(accountUpdateRequest)
}

func (r *accountUpdateRequest) populate(a *model.Account) {
	r.FirstName = a.FirstName
	r.LastName = a.LastName
	r.PhoneNumber = a.PhoneNumber
	r.DateOfBirth = a.DateOfBirth
	r.GenderID = a.GenderID
}

func (r *accountUpdateRequest) bind(c echo.Context, a *model.Account) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	r.FirstName = a.FirstName
	r.LastName = a.LastName
	r.PhoneNumber = a.PhoneNumber
	r.DateOfBirth = a.DateOfBirth
	r.GenderID = a.GenderID

	return nil
}

type verifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}

func (r *verifyEmailRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
