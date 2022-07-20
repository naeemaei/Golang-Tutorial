package main

type Employee struct {
	ID            int
	Name          string
	Type          string
	SalaryOfMonth int
}

func main() {

	var employee1 Employee = Employee{ID: 1, Name: "Mohammad reza", Type: "Manager", SalaryOfMonth: 700}
	var employee2 Employee = Employee{ID: 1, Name: "Peyman", Type: "Manager", SalaryOfMonth: 750}
	salary1 := CalcYearlySalary(employee1)
	salary2 := employee1.CalcYearlySalary()
	salary3 := employee2.CalcYearlySalary()
	println(salary1)
	println(salary2)
	println(salary3)
}

// Function
func CalcYearlySalary(employee Employee) int {
	return employee.SalaryOfMonth * 12
}

// method of Employee struct
func (employee *Employee) CalcYearlySalary() int {
	return employee.SalaryOfMonth * 12
}