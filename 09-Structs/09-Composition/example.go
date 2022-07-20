package main

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate  = 80000
	TaxRate          = 0.09
)

func main() {
	// 
	fullTimeEmployees := []FullTimeEmployee{
		{Employee: Employee{ Id: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee"}, ExtraHours: 40},
		{Employee: Employee{ Id: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini"}, ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{Employee: Employee{ Id: 3, NationalCode: "6563453455", FullName: "Milad Hassani"}, PartTimeHours: 100},
		{Employee: Employee{ Id: 4, NationalCode: "5435435435", FullName: "Maryam Rezaee"}, PartTimeHours: 87},
	}

	shiftEmployees := []ShiftEmployee{
		{Employee: Employee{ Id: 5, NationalCode: "3123123213", FullName: "Shahin"}, ShiftHours: 150},
		{Employee: Employee{ Id: 6, NationalCode: "6363454355", FullName: "Masoud"}, ShiftHours: 168},
	}

	var employees []EmployeeSalaryCalculator = []EmployeeSalaryCalculator{}

	for _, employee := range fullTimeEmployees {
		employees = append(employees, employee)
	}
	
	for _, employee := range partTimeEmployees {
		employees = append(employees, employee)
	}
	
	for _, employee := range shiftEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculate()
		fmt.Printf("\nEmployee (%T): %v\nSalary: %f\nTax: %f\n", employee, employee, salary, tax)
	}
 
}

type EmployeeSalaryCalculator interface {
	SalaryCalculate() (salary float64, tax float64)
}

type Employee struct{
	Id           int
	NationalCode string
	FullName     string
	Unit		 string
}

type FullTimeEmployee struct {
	Employee
	ExtraHours   float64
}

type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

type ShiftEmployee struct {
	Employee
	ShiftHours   float64
}

func (employee FullTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = ShiftSalaryRate * employee.ShiftHours
	tax = salary * TaxRate
	salary += tax
	return
}
 