package sms

import "fmt"

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("soy sms")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "Method: sms"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "sms channel"
}