package main

import (
	"fmt"
	"strings"
)

func main() {

	var notificationType string // "sms,email,push"

	println("Enter notification type: ")

	fmt.Scanln(&notificationType)

	switch {
	case strings.Contains(notificationType, "sms"):
		println("SMS sent")
		fallthrough
	case strings.Contains(notificationType, "email"):
		println("Email sent")
		fallthrough
	case strings.Contains(notificationType, "push"):
		println("Push sent")
	default:
		println("Unknown")
	}

}
