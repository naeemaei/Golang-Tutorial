package main

import "fmt"

type NotificationType string

const (
	Email NotificationType = "Email"
	SMS   NotificationType = "SMS"
)

type Notifier interface {
	Notify(receiver string, message string)
}

type emailNotifier struct {
}

type smsNotifier struct {
}

type User struct {
	Name   string
	Email  string
	Mobile string
	NotificationType
}

type UserService struct{
	Notifier
}

func main() {
	ali := User{Name: "Ali", Email: "ali@ali.com", Mobile: "09121231231", NotificationType: Email}
	mojtaba := User{Name: "Mottaba", Email: "mojtaba@mojtaba.com", Mobile: "09123322114", NotificationType: SMS}
	
	NewNotifier(mojtaba.NotificationType).Notify("Hello",mojtaba.Mobile)
	NewNotifier(ali.NotificationType).Notify("Hello",mojtaba.Email)
}

func (user *User) Notify(message string) {
	switch user.NotificationType {
	case Email:
		fmt.Println("Email sent to", user.Email)
	case SMS:
		fmt.Println("SMS: sent to", user.Mobile)
	}
}

func (notifier emailNotifier) Notify(receiver string,message string) {
	fmt.Println(message, " sent to email", receiver)
}

func (notifier smsNotifier) Notify(receiver string,message string) {
	fmt.Println(message, " sent to mobile", receiver)
}

func NewNotifier(notificationType NotificationType) Notifier {
	switch notificationType {
	case Email:
		return emailNotifier{}
	case SMS:
		return smsNotifier{}
	}
	return nil
}