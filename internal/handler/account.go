package handler

import (
	model "miu-auth-api-v1/internal/model"
	platform "miu-auth-api-v1/internal/platform"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/labstack/echo"
)

func (h *Handler) SignUp(c echo.Context) error {
	var a model.Account
	var p model.Pin
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

	pinResp := generatePin()

	p.ExpiredAt = pinResp.expiredAt
	p.Pin = pinResp.pin
	p.EmailAddress = a.EmailAddress
	p.Purpose = "SignUp"

	if err = h.pinStore.Create(&p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if err = h.sendVerifyEmail(&a, &p); err != nil {
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
	var rt model.RevokedToken
	// req := &revokeTokenRequest{}
	// if err := req.bind(c, &rt); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	// }

	rt.AccountId = int32(getAccountIDFromToken(c))
	rt.ExpiredAt = getExpFromToken(c)
	rt.Jti = getJtiFromToken(c)
	rt.Token = getToken(c)
	rt.RevokedAt = time.Now().UTC()

	ert, err := h.revokedTokenStore.GetByJti(rt.Jti)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if ert != nil {
		return c.JSON(http.StatusOK, newGenericResponse("Logout successfully!!"))
	}

	if err := h.revokedTokenStore.Create(&rt); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}

	return c.JSON(http.StatusCreated, newGenericResponse("Logout successfully!"))
}

func (h *Handler) Change(c echo.Context) error {
	var rt model.RevokedToken
	rt.AccountId = int32(getAccountIDFromToken(c))
	rt.ExpiredAt = getExpFromToken(c)
	rt.Jti = getJtiFromToken(c)
	rt.Token = getToken(c)
	rt.RevokedAt = time.Now().UTC()
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
	ert, er := h.revokedTokenStore.GetByJti(rt.Jti)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if ert != nil {
		return c.JSON(http.StatusOK, newAccountResponse(a))
	}

	if er := h.revokedTokenStore.Create(&rt); er != nil {
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
	req.populateAccount(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	a.LastModifiedAt = time.Now().UTC()
	if err := h.accountStore.Update(a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	return c.JSON(http.StatusOK, newAccountResponse(a))
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	req := &forgotPasswordRequest{}
	var p model.Pin

	if err := req.bind(c, &p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	a, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}
	pinGen := generatePin()
	p.ExpiredAt = pinGen.expiredAt
	p.Pin = pinGen.pin
	p.Purpose = "forgotten"

	if err := h.pinStore.Create(&p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	//TODO Send email
	emErr := h.sendVerifyEmail(a, &p)
	if emErr != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, requestForgotPasswordResponse())
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

func (h *Handler) Verify(c echo.Context) error {
	//req := &verifyEmailRequest{}
	return c.JSON(http.StatusOK, newGenericResponse("Success"))
}
