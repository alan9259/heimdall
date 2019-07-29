package handler

import (
	model "heimdall/internal/model"
	platform "heimdall/internal/platform"
	"net/http"
	"time"

	"github.com/google/uuid"

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

	p, err := h.generatePin(req.EmailAddress, "SignUp")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err = h.sendVerifyEmail(&a, p); err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
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

	a.LastLoginAt = time.Now().UTC()

	if err := h.accountStore.Update(a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}

	return c.JSON(http.StatusOK, newAccountResponse(a))
}

func (h *Handler) Logout(c echo.Context) error {
	rt := newRevokedToken(c)
	ert, err := h.revokedTokenStore.GetByJti(rt.Jti)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if ert != nil {
		return c.JSON(http.StatusOK, newGenericResponse("Logout successfully!!"))
	}

	if err := h.revokedTokenStore.Create(rt); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}

	return c.JSON(http.StatusOK, newGenericResponse("Logout successfully!"))
}

func (h *Handler) Change(c echo.Context) error {
	rt := newRevokedToken(c)
	req := &changeRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	a, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	if !a.CheckPassword(req.OldPassword) {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	hash, err := a.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	a.Password = hash
	if err := h.accountStore.Update(a); err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	ert, err := h.revokedTokenStore.GetByJti(rt.Jti)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if ert != nil {
		return c.JSON(http.StatusOK, newAccountResponse(a))
	}

	if err := h.revokedTokenStore.Create(rt); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	return c.JSON(http.StatusOK, passwordChangeResponse(a))
}

func (h *Handler) GetCurrentAccount(c echo.Context) error {
	a, err := h.accountStore.GetByID(getAccountIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, platform.NotFound())
	}
	return c.JSON(http.StatusOK, newCurrentAccountResponse(a))
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
	req.populateAccount(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	a.LastModifiedAt = time.Now().UTC()
	if err := h.accountStore.Update(a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	req := &forgotPasswordRequest{}

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

	p, err := h.generatePin(req.EmailAddress, "forgot")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.sendVerifyEmail(a, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newForgotPasswordResponse())
}

func (h *Handler) Verify(c echo.Context) error {
	req := &verifyEmailRequest{}

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

	p, err := h.generatePin(req.EmailAddress, "verify")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.sendVerifyEmail(a, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newVerifyEmailResponse())
}

//private func

func getAccountIDFromToken(c echo.Context) uint {
	id, ok := c.Get("accountId").(uint)
	if !ok {
		return 0
	}
	return id
}

func getToken(c echo.Context) string {
	token, ok := c.Get("token").(string)
	if !ok {
		return ""
	}
	return token
}

func getExpFromToken(c echo.Context) time.Time {
	exp, ok := c.Get("tokenExp").(time.Time)
	if !ok {
		return time.Time{}
	}
	return exp
}

func getJtiFromToken(c echo.Context) uuid.UUID {
	jti, ok := c.Get("tokenJti").(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}
	return jti
}

func newRevokedToken(c echo.Context) *model.RevokedToken {
	var rt model.RevokedToken
	rt.AccountId = int32(getAccountIDFromToken(c))
	rt.ExpiredAt = getExpFromToken(c)
	rt.Jti = getJtiFromToken(c)
	rt.Token = getToken(c)
	rt.RevokedAt = time.Now().UTC()

	return &rt
}
