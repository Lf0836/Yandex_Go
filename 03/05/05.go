package main

import (
	"fmt"
)

func main() {
	var a, b, c, n int
	n, _ = fmt.Scan(&a, &b, &c)
	if n != 3 {
		fmt.Print("Некорректный ввод")
	} else if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if a != b && b != c && a != c {
		fmt.Println("Все числа разные")
	} else {
		fmt.Println("Два числа равны")
	}
}
