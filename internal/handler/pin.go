package handler

import (
	"math/rand"
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"

	"time"
)

type pinValidateRequestModel struct {
	emailAddress string
	pin          int
}

type pinValidateResponseModel struct {
	status    int32
	message   string
	token     string
	expiredAt time.Time
}
type pinGenerateResponseModel struct {
	pin       int32
	expiredAt time.Time
}

func generatePin() *pinGenerateResponseModel {
	resp := new(pinGenerateResponseModel)
	t := time.Now()
	resp.pin = int32(rand.Intn(1000000))
	resp.expiredAt = t.AddDate(0, 0, 3)
	return resp
}

func validatePin(a *model.Pin) *pinValidateResponseModel {
	resp := new(pinValidateResponseModel)
	t := time.Now()
	resp.token = platform.GenerateJWTToken(a.Pin)
	resp.expiredAt = t.AddDate(0, 0, 3)
	resp.status = 200
	resp.message = "Success"
	return resp

}
