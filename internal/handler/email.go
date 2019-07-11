package handler

import (
	model "miu-auth-api-v1/internal/model"
	platform "miu-auth-api-v1/internal/platform"
)

func (h *Handler) sendVerifyEmail(a *model.Account) error {
	key, err := h.configStore.GetApiKey("sendgrid API key")
	if err != nil {

	}
	sender := "MIU"
	senderEmail := "alan9259@gmail.com"
	subject := "Thank you for signing up"
	url := "http://miu.com"
	content := "Hi " + a.FirstName + ", <br><br>" + "Thank you for signing up, " +
		"please verify your email address by clicking " + url + "<br><br>" + "MIU"

	platform.SendEmail(
		sender,
		senderEmail,
		subject,
		a.FirstName,
		a.EmailAddress,
		content,
		key.Value)
}
