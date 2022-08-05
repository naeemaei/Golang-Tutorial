package externalServices

import (
	"fmt"

	"github.com/rs/zerolog"
)

type NilNotifyService struct {
}

func (e *NilNotifyService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("nilNotifyService: receiver: %s, message: %s\n", receiver, message)
	logger.Info().Str("notifier", "nilNotifyService").
		Dict("messageInfo", zerolog.Dict().
			Str("receiver", receiver).Str("message", message)).
		Msg("Message sent")

	return nil
}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
