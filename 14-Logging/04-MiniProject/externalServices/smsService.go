package externalServices

import (
	"fmt"
	"notification/entities"

	"github.com/rs/zerolog"
)

type SmsService struct {
}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("sms sent: %v\n", order)
}

func (e *SmsService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}

	fmt.Printf("sms sent: receiver: %s, message: %s\n", receiver, message)

	fmt.Printf("nilNotifyService: receiver: %s, message: %s\n", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("Message sent")

	return nil
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
