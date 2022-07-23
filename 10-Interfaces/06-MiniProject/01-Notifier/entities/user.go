package entities

import "notification/consts"

type User struct {
	Name   string
	Email  string
	Mobile string
	consts.NotificationType
}
