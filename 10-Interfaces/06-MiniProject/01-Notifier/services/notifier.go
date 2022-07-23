package services

import (
	"fmt"
	consts "notification/consts"
)

type Notifier interface {
	Notify(receiver string, message string)
}

type emailService struct {
}

type smsService struct {
}

func (notifier emailService) Notify(receiver string, message string) {
	fmt.Println(message, " sent to email", receiver)
}

func (notifier smsService) Notify(receiver string, message string) {
	fmt.Println(message, " sent to mobile", receiver)
}

func NewNotifier(notificationType consts.NotificationType) Notifier {
	switch notificationType {
	case consts.Email:
		return emailService{}
	case consts.SMS:
		return smsService{}
	}
	return nil
}
