package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name, "\nAge:", p.age, "\nAddress:", p.address)
}
