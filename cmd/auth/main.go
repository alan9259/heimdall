package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func Login(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("i am a fake token")
}

func Logout(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("logged out")
}

func Register(w http.ResponseWriter, req *http.Request) {
	var account Account
	_ = json.NewDecoder(req.Body).Decode(&account)

	json.NewEncoder(w).Encode(account)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("POST")
	router.HandleFunc("/register", Register).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
