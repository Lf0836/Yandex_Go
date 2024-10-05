package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Println("Число положительное")
	}
	if n < 0 {
		fmt.Println("Число отрицательное")
	}
	if n == 0 {
		fmt.Println("Введен ноль")
	}
}
