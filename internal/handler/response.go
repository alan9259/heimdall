package handler

import (
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"
	"time"
)

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

type genericResponse struct {
	Message string `json:"message"`
}

func newGenericResponse(message string) *genericResponse {
	r := new(genericResponse)
	r.Message = message
	return r
}

type changeResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func passwordChangeResponse(a *model.Account) *changeResponse {
	resp := new(changeResponse)
	resp.Token = platform.GenerateJWTToken(a.ID)
	resp.Message = `Your password has been changed successfully.`
	return resp
}

type forgotPasswordResponse struct {
	Message string `json:"message"`
	Status  string `json:"code"`
}

func requestForgotPasswordResponse() *forgotPasswordResponse {
	resp := new(forgotPasswordResponse)
	resp.Message = "We have sent you an email with a reset code for verification."
	resp.Status = "200"
	return resp
}

type pinValidateResponse struct {
	Status    int32
	Message   string
	Token     string
	ExpiredAt time.Time
}

func newPinValidateResponse(id int32, message string) *pinValidateResponse {
	r := new(pinValidateResponse)
	r.ExpiredAt = time.Now().AddDate(0, 0, 3)
	r.Message = message
	r.Status = 200
	r.Token = platform.GenerateJWTToken(id)
	return r
}
