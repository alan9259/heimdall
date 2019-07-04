package model

type Location struct {
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}
