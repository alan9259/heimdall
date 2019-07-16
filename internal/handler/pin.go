package handler

import (
	"math/rand"
	"miu-auth-api-v1/internal/platform"

	"time"
)

type pinModel struct {
	emailAddress string
	pin          int32
	expiredAt    time.Time
	purpose      string
}

type pinValidateRequestModel struct {
	emailAddress string
	pin          int
}

type pinValidateResponseModel struct {
	status    int
	message   string
	token     string
	issuedAt  time.Time
	expiredAt time.Time
}

type pinGenerateRequestModel struct {
	emailAddress string
	pin          int
	expiredAt    time.Time
	purpose      string
}

type pinGenerateResponseModel struct {
	emailAddress string
	pin          int
	expiredAt    time.Time
	purpose      string
}

func generatePin(a *pinModel) *pinGenerateResponseModel {
	resp := new(pinGenerateResponseModel)
	t := time.Now()
	resp.pin = rand.Intn(10000)
	resp.emailAddress = "kennyspeaks@gmail.com"
	resp.expiredAt = t.AddDate(0, 0, 3)
	resp.purpose = "VerifyEmail"
	return resp
}

func validatePin(a *pinModel) *pinValidateResponseModel {
	resp := new(pinValidateResponseModel)
	t := time.Now()
	resp.token = platform.GenerateJWTToken(a.pin)
	resp.expiredAt = t.AddDate(0, 0, 3)
	resp.issuedAt = t
	resp.status = 200
	resp.message = "Success"
	return resp

}
