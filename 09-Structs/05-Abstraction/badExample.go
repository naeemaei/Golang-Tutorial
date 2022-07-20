package main

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate = 80000
	TaxRate          = 0.09
)

func main() {
	employees := []Employee{
		{Id: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee", Type: "FullTime", Hours: 40},
		{Id: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini", Type: "PartTime", Hours: 120},
		{Id: 3, NationalCode: "6542541458", FullName: "Reza ", Type: "FullTime", Hours: 15.5},
		{Id: 4, NationalCode: "1368245856", FullName: "Mahtab", Type: "Shift", Hours: 160},
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculator(float64(employee.Hours))
		fmt.Printf("Employee: %v\nSalary: %f\nTax: %f\n", employee, salary, tax)
	}
}

type Employee struct {
	Id           int
	NationalCode string
	FullName     string
	Type         string
	Hours		 float64
}

func (employee Employee) FullTimeSalaryCalculator(extraHours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) PartTimeSalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) ShiftSalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = ShiftSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	if employee.Type == "FullTime" {
		return employee.FullTimeSalaryCalculator(hours)
	} else if employee.Type == "PartTime" {
		return employee.PartTimeSalaryCalculator(hours)
	} else{
		return employee.ShiftSalaryCalculator(hours)
	}
}
