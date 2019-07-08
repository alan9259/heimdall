package main

import (
	"miu-auth-api-v1/internal/handler"
	"miu-auth-api-v1/internal/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := platform.New()
	//db.AutoMigrate(d)

	as := platform.NewAccountStore(d)
	ls := platform.NewLocationStore(d)
	h := handler.NewHandler(as, ls)

	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}

/*import (
	"encoding/json"
	"log"
	model "miu-auth-api-v1/internal/model"
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

}*/
