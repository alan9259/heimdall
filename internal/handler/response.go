package handler

import (
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"
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
	Message string `json: message`
	Token   string `json: "token"`
}

func passwordChangeResponse(a *model.Account) *changeResponse {
	resp := new(changeResponse)
	resp.Token = platform.GenerateJWTToken(a.ID)
	resp.Message = `Your password been changed successfully.`
	return resp
}
