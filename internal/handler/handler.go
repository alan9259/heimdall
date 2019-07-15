package handler

import (
	account "miu-auth-api-v1/internal/account"
	config "miu-auth-api-v1/internal/config"
	location "miu-auth-api-v1/internal/location"
	revokedToken "miu-auth-api-v1/internal/revokedToken"
)

type Handler struct {
	accountStore      account.Store
	locationStore     location.Store
	configStore       config.Store
	revokedTokenStore revokedToken.Store
}

func NewHandler(
	as account.Store,
	ls location.Store,
	cs config.Store,
	rts revokedToken.Store) *Handler {
	return &Handler{
		accountStore:      as,
		locationStore:     ls,
		configStore:       cs,
		revokedTokenStore: rts,
	}
}
