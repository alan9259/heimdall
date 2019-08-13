package handler

import (
	"heimdall/internal/interface/pin"
	model "heimdall/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ps pin.Store
)

func TestGeneratePinSuccess(t *testing.T) { //Happy path
	tearDown()
	setup()
	res, err := h.generatePin("declan.grogan8@mail.dcu.ie", "verify")
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Pin)
	assert.Equal(t, res.Purpose, "verify")
}

func TestCheckDupicates(t *testing.T) {
	tearDown()
	setup()
	setUpMockData(t)
	pin := model.Pin{
		EmailAddress: "declan.grogan8@mail.dcu.ie",
		Pin:          123412,
		Purpose:      "verify",
		ExpiredAt:    time.Now().AddDate(0, 0, 3),
	}
	h.checkDuplicates(&pin)
	pins, err := h.pinStore.GetCurrentPins(pin.EmailAddress)
	if err != nil {
		t.Fatal("There was an error getting current pins")
	}
	assert.NotNil(t, pins)
	assert.Equal(t, pin.Purpose, (*pins)[0].Purpose)
	assert.NotEqual(t, pin.Pin, (*pins)[0].Pin)
}

func TestCompareToCurrentPins(t *testing.T) {
	tearDown()
	setup()
	setUpMockData(t)
	pin := model.Pin{
		EmailAddress: "declan.grogan8@mail.dcu.ie",
		Pin:          123412,
		Purpose:      "verify",
		ExpiredAt:    time.Now().AddDate(0, 0, 3),
	}
	if err := h.compareToCurrentPins(&pin); err != nil {
		t.Logf("An error occurred comparing the generated pin to existing pins")
	}
	pins, err := h.pinStore.GetCurrentPins(pin.EmailAddress)
	if err != nil {
		t.Fatal("There was an error getting current pins")
	}
	assert.Nil(t, pins)
}

func setUpMockData(t *testing.T) {
	pin := model.Pin{
		EmailAddress: "declan.grogan8@mail.dcu.ie",
		Pin:          123412,
		Purpose:      "verify",
		ExpiredAt:    time.Now().AddDate(0, 0, 3),
	}

	if err := h.pinStore.Create(&pin); err != nil {
		t.Error("Creation of mock pin has failed.")
	}
	t.Logf("Created a pin %d", pin.Pin)
}
