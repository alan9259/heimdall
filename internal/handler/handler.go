package handler

import (
	account "miu-auth-api-v1/internal/account"
	location "miu-auth-api-v1/internal/location"
)

type Handler struct {
	accountStore  account.Store
	locationStore location.Store
}

func NewHandler(as account.Store, ls location.Store) *Handler {
	return &Handler{
		accountStore:  as,
		locationStore: ls,
	}
}

func NewMockHandler() *Handler {
	return &Handler{}
}
