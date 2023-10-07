// main -> orderService -> emailService || smsService
package main

import (
	"log"
	"notification/core"
	"notification/entities"
	"notification/services"
)

var logger = core.NewFileLogger()

func main() {
	log.Logger = logger
	order1 := entities.Order{
		ID:               1,
		UserFullName:     "John Doe",
		UserId:           "09115252524",
		Price:            float64(100),
		Status:           true,
		NotificationType: entities.Email,
	}
	order2 := entities.Order{
		ID:               2,
		UserFullName:     "John Mac",
		UserId:           "09123332221",
		Price:            float64(220),
		Status:           true,
		NotificationType: entities.Sms,
	}

	order3 := entities.Order{
		ID:               3,
		UserFullName:     "Reza Mac",
		UserId:           "",
		Price:            float64(210),
		Status:           true,
		NotificationType: "",
	}

	orderService := services.NewOrderService()
	err, _ := orderService.CreateOrder(&order1)
	if err != nil {
		logger.Error().Interface("entityInfo", order1).Err(err).Msg("Order 1 is not valid")
	}
	err, _ = orderService.CreateOrder(&order2)

	if err != nil {
		logger.Error().Interface("entityInfo", order2).Err(err).Msg("Order 1 is not valid")
	}

	err, _ = orderService.CreateOrder(&order3)
	if err != nil {
		logger.Error().Interface("entityInfo", order3).Err(err).Msg("Order 1 is not valid")
	}
}
