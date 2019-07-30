package email

type Service interface {
	SendEmail(
		sender string,
		senderEmail string,
		subject string,
		receiver string,
		receiverEmail string,
		emailContent string) error
}
