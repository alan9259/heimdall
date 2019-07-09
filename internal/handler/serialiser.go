package handler

import (
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"

	"github.com/labstack/echo"
)

type registerRequest struct {
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

type accountResponse struct {
	EmailAddress string `json:"email_address"`
	Token        string `json:"token"`
}

func newAccountResponse(a *model.Account) *accountResponse {
	r := new(accountResponse)
	r.EmailAddress = a.EmailAddress
	r.Token = platform.GenerateJWTToken(a.ID)
	return r
}
