package handler

import (
	"log"
	model "heimdall/internal/model"
	platform "heimdall/internal/platform"
	"strconv"
)

func (h *Handler) sendVerifyEmail(a *model.Account, p *model.Pin) error {
	key, err := h.configStore.GetApiKey("sendgridApikey")

	if err != nil {
		log.Println("Error: Failed to fetch sendgrid api key from database")
		return err
	}

	sender := "MIU"
	senderEmail := "alan9259@gmail.com"
	subject := "Thank you for signing up"
	content := ""
	if p.Purpose == "SignUp" {
		content = "Hi " + a.FirstName + ", <br><br>" + "Thank you for signing up, " +
			"please verify your email address by entering the verification code: " + strconv.Itoa(int(p.Pin)) + " in your app." + "<br><br>" + "MIU"
	} else {
		content = "Hi " + a.FirstName + ", <br><br>" + "You've forgotten your password! " +
			"You will find the reset password screen after entering the verification code: " + strconv.Itoa(int(p.Pin)) + " in your app." + "<br><br>" + "MIU"
	}

	err = platform.SendEmail(
		sender,
		senderEmail,
		subject,
		a.FirstName,
		a.EmailAddress,
		content,
		key.Value)

	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	return nil
}
