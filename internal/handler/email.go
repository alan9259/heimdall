package handler

import (
	model "heimdall/internal/model"
	"log"
	"strconv"
)

func (h *Handler) sendVerifyEmail(a *model.Account, p *model.Pin) error {
	sender := "MIU"
	senderEmail := "alan9259@gmail.com"
	subject := "Thank you for signing up"
	content := ""
	strPin := strconv.Itoa(int(p.Pin))
	if p.Purpose == "verify" {
		content = "Hi " + a.FirstName + ", <br><br>" + "Thank you for signing up, " +
			"please verify your email address by entering the verification code: " + padZeros(strPin) + " in your app." + "<br><br>" + "MIU"
	} else {
		content = "Hi " + a.FirstName + ", <br><br>" + "You've forgotten your password! " +
			"You will find the reset password screen after entering the verification code: " + padZeros(strPin) + " in your app." + "<br><br>" + "MIU"
	}

	err := h.emailService.SendEmail(
		sender,
		senderEmail,
		subject,
		a.FirstName,
		a.EmailAddress,
		content)

	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	return nil
}

func padZeros(pin string) (result string) {
	var pinStr = pin
	if len(pinStr) == 6 {
		return pinStr
	}
	pinStr = "0" + pinStr
	return padZeros(pinStr)
}
