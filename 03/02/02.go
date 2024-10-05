package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n > m {
		fmt.Println("Первое число больше второго")
	}
	if n < m {
		fmt.Println("Второе число больше первого")
	}
	if n == m {
		fmt.Println("Числа равны")
	}
}
