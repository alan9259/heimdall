package platform

import (
	"errors"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	client *sendgrid.Client
	isMock bool
}

func NewEmailService(key string, isMock bool) *EmailService {
	if isMock {
		return &EmailService{
			client: nil,
			isMock: isMock,
		}
	} else {
		return &EmailService{
			client: sendgrid.NewSendClient(key),
			isMock: isMock,
		}
	}
}

func (es *EmailService) SendEmail(
	sender string,
	senderEmail string,
	subject string,
	receiver string,
	receiverEmail string,
	emailContent string) error {

	if len(sender) == 0 {
		log.Println("You must specify a sender name")
		return errors.New("You must specify a sender name")
	}

	if len(senderEmail) == 0 {
		log.Println("You must specify a valid sender email address")
		return errors.New("You must specify a valid sender email address")
	}

	if len(subject) == 0 {
		log.Println("You must specify a subject for this email")
		return errors.New("You must specify a subject for this email")
	}

	if len(receiver) == 0 {
		log.Println("You must specify a receiver name")
		return errors.New("You must specify a receiver name")
	}

	if len(receiverEmail) == 0 {
		log.Println("You must specify a valid receiver email address")
		return errors.New("You must specify a valid receiver email address")
	}

	if len(emailContent) == 0 {
		log.Println("You must specify a content for this email")
		return errors.New("You must specify a content for this email")
	}

	from := mail.NewEmail(sender, senderEmail)
	emailTitle := subject
	to := mail.NewEmail(receiver, receiverEmail)
	plainTextContent := emailContent
	htmlContent := "<strong>" + emailContent + "</strong>"
	message := mail.NewSingleEmail(from, emailTitle, to, plainTextContent, htmlContent)

	//client := sendgrid.NewSendClient(key) // key was passed in as env variable
	if !es.isMock {

		response, err := es.client.Send(message)

		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println(response.StatusCode)
			log.Println(response.Body)
			log.Println(response.Headers)
		}

	}

	return nil
}
