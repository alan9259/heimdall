package handler

import (
	"heimdall/internal/interface/account"
	"heimdall/internal/interface/config"
	"heimdall/internal/interface/email"
	"heimdall/internal/interface/location"
	"heimdall/internal/interface/pin"
	"heimdall/internal/interface/revokedToken"
)

type Handler struct {
	accountStore      account.Store
	locationStore     location.Store
	configStore       config.Store
	pinStore          pin.Store
	revokedTokenStore revokedToken.Store
	emailService      email.Service
}

func NewHandler(
	as account.Store,
	ls location.Store,
	cs config.Store,
	ps pin.Store,
	rts revokedToken.Store,
	es email.Service) *Handler {
	return &Handler{
		accountStore:      as,
		locationStore:     ls,
		configStore:       cs,
		pinStore:          ps,
		revokedTokenStore: rts,
		emailService:      es,
	}
}
