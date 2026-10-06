package notifications

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

type SMSNotifier struct {
	ProneNumber string
}

func (email EmailNotifier) Notify(message string) {
	fmt.Printf("\nОтправляю email на %v: '%v'\n", email.EmailAddress, message)
}

func (sms SMSNotifier) Notify(message string) {
	fmt.Printf("\nОтправляю SMS на номер %v: '%v'\n", sms.ProneNumber, message)
}
