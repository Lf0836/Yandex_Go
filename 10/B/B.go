package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Print("Employee: ", e.name, "\nPosition: ", e.position, "\nTotal Salary: ")
	fmt.Printf("%.2f", e.salary+e.bonus)
}
