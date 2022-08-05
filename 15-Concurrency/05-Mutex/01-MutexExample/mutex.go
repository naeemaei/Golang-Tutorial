package main

import (
	"sync"
	"sync/atomic"
)

type Employee struct {
	Id     int
	Salary int64
}

func MutexAccountDecrease() {
	println("Start of MutexAccountDecrease")
	println("Write 25_500_000_000 - 25_000_000_000 = 500_000_000")
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	var money int64 = 25_500_000_000
	println(money)
	employeeSalaryList := []Employee{}
	wg.Add(5000)
	for i := 0; i < 5_000; i++ {
		employeeSalaryList = append(employeeSalaryList, Employee{Id: i, Salary: GetRandomNumber()})
	}

	for _, employee := range employeeSalaryList {
		go func(employee Employee) {
			defer wg.Done()
			if employee.Salary < money {
				mx.Lock()
				money -= employee.Salary
				mx.Unlock()
			}
		}(employee)

	}

	wg.Wait()

	println(money)

}

func AtomicAccountDecrease() {
	println("Start of AtomicAccountDecrease")
	println("Write 25_500_000_000 - 25_000_000_000 = 500_000_000")
	wg := sync.WaitGroup{}

	var money int64 = 25_500_000_000
	println(money)
	employeeSalaryList := []Employee{}
	wg.Add(5000)
	for i := 0; i < 5_000; i++ {
		employeeSalaryList = append(employeeSalaryList, Employee{Id: i, Salary: GetRandomNumber()})
	}

	for _, employee := range employeeSalaryList {
		go func(employee Employee) {
			defer wg.Done()
			if employee.Salary < money {
				atomic.AddInt64(&money, -employee.Salary)
			}
		}(employee)

	}

	wg.Wait()

	println(money)

}

func BadAccountDecrease() {
	println("Start of BadAccountDecrease")
	println("Write 25_500_000_000 - 25_000_000_000 = 500_000_000")
	wg := sync.WaitGroup{}

	var money int64 = 25_500_000_000
	println(money)
	employeeSalaryList := []Employee{}
	wg.Add(5000)
	for i := 0; i < 5_000; i++ {
		employeeSalaryList = append(employeeSalaryList, Employee{Id: i, Salary: GetRandomNumber()})
	}

	for _, employee := range employeeSalaryList {
		go func(employee Employee) {
			defer wg.Done()
			if employee.Salary < money {
				money -= employee.Salary
			}
		}(employee)

	}

	wg.Wait()

	println(money)

}

func GetRandomNumber() int64 {
	return 5_000_000
}
