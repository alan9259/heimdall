package handler

import (
	model "miu-auth-api-v1/internal/model"
	platform "miu-auth-api-v1/internal/platform"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) SignUp(c echo.Context) error {
	var a model.Account
	req := &registerRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewError(err))
	}
	if err := h.accountStore.Create(&a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewError(err))
	}
	return c.JSON(http.StatusCreated, newAccountResponse(&a))
}

func (h *Handler) Login(c echo.Context) error {
	req := &loginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewError(err))
	}
	a, err := h.accountStore.GetByEmail(req.Account.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	if !a.CheckPassword(req.Account.Password) {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newAccountResponse(a))
}
