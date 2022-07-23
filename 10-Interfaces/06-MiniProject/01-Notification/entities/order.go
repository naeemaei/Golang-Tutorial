package entities

type Order struct {
	ID               int64
	UserFullName     string
	UserId           string
	Price            float64
	Status           bool
	NotificationType NotificationType
}
