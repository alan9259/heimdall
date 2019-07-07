package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(Sender string, SenderEmail string, Subject string, Receiver string, ReceiverEmail string, EmailContent string) {
	from := mail.NewEmail(Sender, SenderEmail)
	subject := Subject
	to := mail.NewEmail(Receiver, ReceiverEmail)
	plainTextContent := EmailContent
	htmlContent := "<strong>" + EmailContent + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")) // key was passed in as env variable

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
