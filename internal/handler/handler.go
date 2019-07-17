package handler

import (
	account "miu-auth-api-v1/internal/account"
	config "miu-auth-api-v1/internal/config"
	location "miu-auth-api-v1/internal/location"
	"miu-auth-api-v1/internal/pin"
	revokedToken "miu-auth-api-v1/internal/revokedToken"
)

type Handler struct {
	accountStore      account.Store
	locationStore     location.Store
	configStore       config.Store
	pinStore          pin.Store
	revokedTokenStore revokedToken.Store
}

func NewHandler(
	as account.Store,
	ls location.Store,
	cs config.Store,
	ps pin.Store,
	rts revokedToken.Store) *Handler {
	return &Handler{
		accountStore:      as,
		locationStore:     ls,
		configStore:       cs,
		pinStore:          ps,
		revokedTokenStore: rts,
	}
}
