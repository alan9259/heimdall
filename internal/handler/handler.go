package handler

import (
	account "heimdall/internal/account"
	config "heimdall/internal/config"
	location "heimdall/internal/location"
	"heimdall/internal/pin"
	revokedToken "heimdall/internal/revokedToken"
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
