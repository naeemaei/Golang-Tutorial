package models

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}
