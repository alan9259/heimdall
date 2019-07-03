package main

import (
	"encoding/json"
	"log"
	model "miu-auth-api-v1/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("i am a fake token")
}

func Logout(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("logged out")
}

func Register(w http.ResponseWriter, req *http.Request) {
	var account model.Account
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
