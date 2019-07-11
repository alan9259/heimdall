package handler

import (
	account "miu-auth-api-v1/internal/account"
	config "miu-auth-api-v1/internal/config"
	location "miu-auth-api-v1/internal/location"
)

type Handler struct {
	accountStore  account.Store
	locationStore location.Store
	configStore   config.Store
}

func NewHandler(as account.Store, ls location.Store, cs config.Store) *Handler {
	return &Handler{
		accountStore:  as,
		locationStore: ls,
		configStore:   cs,
	}
}
