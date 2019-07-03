package model

type Account struct {
	ID           string    `json:"id,omitempty"`
	EmailAddress string    `json:"email_address,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Location     *Location `json:"location,omitempty"`
}

type Location struct {
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}
