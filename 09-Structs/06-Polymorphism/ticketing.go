package main

import "fmt"

type BusTicket struct {
	Id                int
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	BusType           string
	PassengerName     string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}

type FlightTicket struct {
	Id               int
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	AirplaneType     string
	PassengerName    string
	PassportId       string
	PassengerType    string
	Price            int
}

func main() {

	busTicket := BusTicket{
		Id:                1,
		DepartureCity:     "Tehran",
		ArrivalCity:       "Mashhad",
		DepartureTime:     "12:00",
		BusType:           "Bus",
		PassengerName:     "Reza kamali",
		DepartureTerminal: "Terminal 1",
		ArrivalTerminal:   "Terminal 2",
		NationalCode:      "123456789",
		Price:             100,
	}

	flightTicket := FlightTicket{
		Id:               2,
		DepartureAirport: "Tehran",
		ArrivalAirport:   "London",
		DepartureTime:    "12:00",
		ArrivalTime:      "23:00",
		AirplaneType:     "Airbus",
		PassengerName:    "Peyman Hassani",
		PassportId:       "312321414",
		PassengerType:    "Adult",
		Price:            1890,
	}

	printer := []TicketPrinter{busTicket, flightTicket}

	for _, printer := range printer {
		printer.PrintTicket()
	}

 

}

type TicketPrinter interface {
	PrintTicket()
}

func (ticket BusTicket) PrintTicket() {
	fmt.Printf("Bus Ticket:\n ID: %d\n DepartureCity : %s\n ArrivalCity : %s\n DepartureTime : %s\n", ticket.Id, ticket.DepartureCity, ticket.ArrivalCity, ticket.DepartureTime)
	fmt.Printf("BusType : %s\n PassengerName : %s\n DepartureTerminal : %s\n ArrivalTerminal : %s\n", ticket.BusType, ticket.PassengerName, ticket.DepartureTerminal, ticket.ArrivalTerminal)
	fmt.Printf("NationalCode : %s\n Price : %d\n", ticket.NationalCode, ticket.Price)
}

func (ticket FlightTicket) PrintTicket() {
	fmt.Printf("Flight Ticket:\n ID: %d\n DepartureAirport : %s\n ArrivalAirport : %s\n DepartureTime : %s\n ArrivalTime : %s\n", ticket.Id, ticket.DepartureAirport, ticket.ArrivalAirport, ticket.DepartureTime,ticket.ArrivalTime)
	fmt.Printf("AirplaneType : %s\n PassengerName : %s\n PassportId : %s\n PassengerType : %s\n", ticket.AirplaneType, ticket.PassengerName, ticket.PassportId, ticket.PassengerType)
	fmt.Printf("Price : %d\n", ticket.Price)
}