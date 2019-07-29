package handler

import (
	"math/rand"
	"miu-auth-api-v1/internal/model"
	"miu-auth-api-v1/internal/platform"
	"net/http"
	"strconv"
	"strings"

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
	var p model.Pin
	p.ExpiredAt = time.Now().AddDate(0, 0, 3)
	p.Pin = int32(rand.Intn(1000000)) //add a minimum
	p.Purpose = purpose
	p.EmailAddress = email
	s := strconv.Itoa(int(p.Pin))

	if err := h.pinStore.Create(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
func padZeros(int p) (p int) {
	var ap = []int32{P}
	s := strings.Split(str, "")
	l := len(str)
	for l < 6 {
		s = append(ap, str)
		str = strings.Join(s, "")
		l = len(str)
	}
}

/*
str := "1234"
   var ap = []string{"0"}
   s := strings.Split(str, "")
   l := len(str)
   for l < 6 {
      s = append(ap, str)
      str = strings.Join(s, "")
      l = len(str)
   }
*/
