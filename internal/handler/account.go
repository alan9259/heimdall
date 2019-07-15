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
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}

	//check if an existing account has taken the same email address
	ea, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if ea != nil {
		return c.JSON(http.StatusOK, platform.AlreadyRegistered())
	}

	if err := h.accountStore.Create(&a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}

	err = h.sendVerifyEmail(&a)

	if err != nil {
		return c.JSON(http.StatusCreated, platform.NewHttpError(err))
	}

	return c.JSON(http.StatusCreated, newAccountResponse(&a))
}

func (h *Handler) Login(c echo.Context) error {
	req := &loginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	a, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	if !a.CheckPassword(req.Password) {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newAccountResponse(a))
}

func (h *Handler) Reset(c echo.Context) error {
	req := &resetRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewError(err))
	}
	a, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, platform.NewError(err))
	}
	if !a.CheckPassword(req.OldPassword) {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	hash, err := a.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	a.Password = hash
	if err := h.accountStore.UpdateAccountDetails(a); err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewError(err))
	}
	return c.JSON(http.StatusOK, passwordResetResponse(a))
func (h *Handler) GetCurrentAccount(c echo.Context) error {
	a, err := h.accountStore.GetByID(getAccountIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, platform.NotFound())
	}
	return c.JSON(http.StatusOK, newAccountResponse(a))
}

func (h *Handler) UpdateAccount(c echo.Context) error {
	a, err := h.accountStore.GetByID(getAccountIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, platform.NotFound())
	}
	req := newAccountUpdateRequest()
	req.populate(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	if err := h.accountStore.Update(a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	return c.JSON(http.StatusOK, newAccountResponse(a))
}

//private func

func getAccountIDFromToken(c echo.Context) uint {
	id, ok := c.Get("account").(uint)
	if !ok {
		return 0
	}
	return id
}

func (h *Handler) Verify(c echo.Context) error {
	//req := &verifyEmailRequest{}
	return c.JSON(http.StatusOK, newGenericResponse("Success"))
}
