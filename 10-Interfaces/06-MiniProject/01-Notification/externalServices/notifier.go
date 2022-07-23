package externalServices

import "notification/entities"

type Notifier interface {
	SendNotify(receiver string, message string)
}

func NewNotifier(notificationType entities.NotificationType) Notifier {
	switch notificationType {
	case entities.Sms:
		return NewSmsService()
	case entities.Email:
		return NewEmailService()
	default:
		return NewNilNotifyService()
	}
}
