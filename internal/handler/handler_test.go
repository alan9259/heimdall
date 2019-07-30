package handler

import (
	"log"
	"os"
	"testing"
	"time"

	"encoding/json"

	"heimdall/internal/interface/account"
	"heimdall/internal/model"
	"heimdall/internal/platform"
	"heimdall/internal/router"
	"heimdall/internal/store"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var (
	d  *gorm.DB
	as account.Store
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func authHeader(token string) string {
	return "Bearer " + token
}

func setup() {
	d = platform.InitTestDB()
	platform.AutoMigrate(d)
	as := store.NewAccountStore(d)
	ls := store.NewLocationStore(d)
	cs := store.NewConfigStore(d)
	ps := store.NewPinStore(d)
	rts := store.NewRevokedTokenStore(d)

	es := platform.NewEmailService("", true)
	h = NewHandler(as, ls, cs, ps, rts, es)
	e = router.New()
	//setupMockData()
}

func tearDown() {
	_ = d.Close()
	if err := platform.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func responseMap(b []byte) map[string]interface{} {
	var m interface{}
	json.Unmarshal(b, &m)
	return m.(map[string]interface{})
}

func setupMockData() error {
	dob, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")

	if err != nil {
		println("err")
	}

	a1 := model.Account{
		EmailAddress: "test123@miu.com",
		FirstName:    "testF",
		LastName:     "testL",
		Password:     "1234567",
		PhoneNumber:  "0987654321",
		DateOfBirth:  dob,
	}
	a1.Password, _ = a1.HashPassword("secret")
	if err := as.Create(&a1); err != nil {
		return err
	}

	return nil
}
