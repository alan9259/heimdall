package handler

import (
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"

	"github.com/labstack/echo"
)

type registerRequest struct {
	Account struct {
		EmailAddress string `json:"email_address" validate:"required,email_address"`
		Password     string `json:"password" validate:"required"`
	} `json:"account"`
}

func (r *registerRequest) bind(c echo.Context, a *model.Account) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.EmailAddress = r.Account.EmailAddress
	h, err := a.HashPassword(r.Account.Password)
	if err != nil {
		return err
	}
	a.Password = h
	return nil
}

type loginRequest struct {
	Account struct {
		EmailAddress string `json:"email_address" validate:"required,email"`
		Password     string `json:"password" validate:"required"`
	} `json:"account"`
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

type accountResponse struct {
	User struct {
		EmailAddress string `json:"email_address"`
		Token        string `json:"token"`
	} `json:"account"`
}

func newAccountResponse(a *model.Account) *accountResponse {
	r := new(accountResponse)
	r.User.EmailAddress = a.EmailAddress
	r.User.Token = platform.GenerateJWTToken(a.ID)
	return r
}
