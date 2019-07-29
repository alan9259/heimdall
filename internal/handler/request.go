package handler

import (
	"heimdall/internal/model"
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

func (r *accountUpdateRequest) populateAccount(a *model.Account) {
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

	a.FirstName = r.FirstName
	a.LastName = r.LastName
	a.PhoneNumber = r.PhoneNumber
	a.DateOfBirth = r.DateOfBirth
	a.GenderID = r.GenderID

	return nil
}

type changeRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	OldPassword  string `json:"old_password" validate:"required"`
	NewPassword  string `json:"new_password" validate:"required"`
}

func (r *changeRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type forgotPasswordRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

func (r *forgotPasswordRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type verifyEmailRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
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

type pinValidateRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Pin          int32  `json:"pin" validate:"required"`
}

func (r *pinValidateRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
