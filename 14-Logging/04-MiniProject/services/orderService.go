package services

import (
	"errors"
	"fmt"
	"notification/core"
	"notification/entities"
	"notification/externalServices"
)

var logger = core.NewFileLogger()

type OrderService struct {
	externalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) (error, *entities.Order) {
	if !order.Status {
		return errors.New("order is not valid"), order
	}

	if order.Price < 150 {
		return errors.New("order price is not valid"), order
	}

	fmt.Printf("Order created: %v\n", order)

	logger.Info().Interface("Order", order).Msg("Order created:")

	orderService.Notifier = externalServices.NewNotifier(order.NotificationType)

	logger.Info().Msgf("Notifier founded: %T", orderService.Notifier)

	err := orderService.SendNotify(order.UserId, "Order created")

	return err, order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
