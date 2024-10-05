package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n > 0 && m > 0 {
		fmt.Println("Оба числа положительные")
	}
	if n < 0 && m < 0 {
		fmt.Println("Оба числа отрицательные")
	}
	if n > 0 && m < 0 || n < 0 && m > 0 {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
	if n == 0 || m == 0 {
		fmt.Println("Одно из чисел равно нулю")
	}
}
