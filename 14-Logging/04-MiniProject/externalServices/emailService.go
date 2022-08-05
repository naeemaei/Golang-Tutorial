package externalServices

import (
	"fmt"
	"notification/core"
	"notification/entities"

	"github.com/rs/zerolog"
)

var logger = core.NewFileLogger()

type EmailService struct {
}

func (e *EmailService) SendMessage(order *entities.Order) {
	fmt.Printf("Email sent: %v\n", order)
}

func (e *EmailService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}

	fmt.Printf("Email sent: receiver: %s, message: %s\n", receiver, message)

	fmt.Printf("nilNotifyService: receiver: %s, message: %s\n", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("Message sent")

	return nil
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
