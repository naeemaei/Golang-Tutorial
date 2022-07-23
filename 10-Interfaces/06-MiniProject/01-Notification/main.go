// main -> orderService -> emailService || smsService
package main

import (
	"notification/entities"
	"notification/services"
)

func main() {

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
		UserId:           "09991252525",
		Price:            float64(210),
		Status:           true,
		NotificationType: "",
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&order1)

	orderService.CreateOrder(&order2)

	orderService.CreateOrder(&order3)
}
