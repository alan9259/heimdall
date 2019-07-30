package main

import (
	"heimdall/internal/handler"
	"heimdall/internal/platform"
	"heimdall/internal/router"
	"heimdall/internal/store"
	"log"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := platform.New()
	platform.AutoMigrate(d)

	as := store.NewAccountStore(d)
	ls := store.NewLocationStore(d)
	cs := store.NewConfigStore(d)
	ps := store.NewPinStore(d)
	rts := store.NewRevokedTokenStore(d)

	key, err := cs.GetApiKey("sendgridApikey")

	if err != nil {
		log.Println("Error: Failed to fetch sendgrid api key from database")
		panic(err)
	}

	es := platform.NewEmailService(key.Value, false)

	h := handler.NewHandler(as, ls, cs, ps, rts, es)

	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}

/*import (
	"encoding/json"
	"log"
	model "heimdall/internal/model"
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
