package main

import (
	"notification/consts"
	"notification/entities"
	"notification/services"
)

func main() {
	ali := entities.User{Name: "Ali", Email: "ali@ali.com", Mobile: "09121231231", NotificationType: consts.Email}
	mojtaba := entities.User{Name: "Mottaba", Email: "mojtaba@mojtaba.com", Mobile: "09123322114", NotificationType: consts.SMS}

	services.NewNotifier(mojtaba.NotificationType).Notify("Hello", mojtaba.Mobile)
	services.NewNotifier(ali.NotificationType).Notify("Hello", mojtaba.Email)
}
