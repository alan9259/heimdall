package platform

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(Sender string, SenderEmail string, Subject string, Receiver string, ReceiverEmail string, EmailContent string) {
	if len(Sender) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a sender name")
	}

	if len(SenderEmail) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a valid sender email address")
	}

	if len(Subject) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a subject for this email")
	}

	if len(Receiver) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a receiver name")
	}

	if len(ReceiverEmail) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a valid receiver email address")
	}

	if len(EmailContent) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a content for this email")
	}

	from := mail.NewEmail(Sender, SenderEmail)
	subject := Subject
	to := mail.NewEmail(Receiver, ReceiverEmail)
	plainTextContent := EmailContent
	htmlContent := "<strong>" + EmailContent + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	key := ""
	client := sendgrid.NewSendClient(key) // key was passed in as env variable

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
