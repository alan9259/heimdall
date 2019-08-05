package handler

import (
	"fmt"
	"heimdall/internal/model"
	"heimdall/internal/platform"
	"math/rand"
	"net/http"

	"time"

	"github.com/labstack/echo"
)

type pinGenerateResponse struct {
	pin       int32
	expiredAt time.Time
}

func (h *Handler) ValidatePin(c echo.Context) error {
	req := &pinValidateRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, platform.NewHttpError(err))
	}
	p, err := h.pinStore.GetByCompositeKey(req.EmailAddress, req.Pin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if p == nil {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}

	if p.ExpiredAt.Before(time.Now()) {
		return c.JSON(http.StatusForbidden, newGenericResponse("Pin is expired"))
	}

	a, err := h.accountStore.GetByEmail(req.EmailAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
	}
	if a == nil {
		return c.JSON(http.StatusForbidden, platform.AccessForbidden())
	}

	if a.VerifiedAt.IsZero() {
		a.VerifiedAt = time.Now()
		if err := h.accountStore.Update(a); err != nil {
			return c.JSON(http.StatusInternalServerError, platform.NewHttpError(err))
		}
	}

	return c.JSON(http.StatusAccepted, newPinValidateResponse(a.ID, "Pin validated successfully"))
}

func (h *Handler) generatePin(email string, purpose string) (*model.Pin, error) {
	p := genPin(email, purpose)
	var pinExists = true
	for pinExists == true {
		prev, err := h.pinStore.GetByCompositeKey(p.EmailAddress, p.Pin)
		if err != nil {
			return nil, err
		}
		if prev != nil { //we randomly got a duplicate pin and email combo.
			fmt.Println("Oh boy we found a duplicate.")
			p = genPin(email, purpose)
			if prev.Purpose == p.Purpose {
				err = h.pinStore.RemovePin(prev.EmailAddress, prev.Pin)
				if err != nil {
					return nil, err
				}
			}
		}
		if prev == nil {
			pinExists = false
		}
	}
	//TODO check for older pin with same purpose and remove those older pins.
	if err := h.pinStore.Create(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func genPin(email string, purpose string) model.Pin {
	var p model.Pin
	p.ExpiredAt = time.Now().AddDate(0, 0, 3)
	p.Pin = int32(rand.Intn(1000000))
	p.Purpose = purpose
	p.EmailAddress = email
	return p
}
