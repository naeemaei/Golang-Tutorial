package externalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("sms sent: %v\n", order)
}

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("sms sent: receiver: %s, message: %s\n", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
