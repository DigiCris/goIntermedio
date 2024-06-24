package email

import "fmt"

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("soy email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Method: email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "email channel"
}