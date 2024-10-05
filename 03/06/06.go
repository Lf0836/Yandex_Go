package main

import (
	"fmt"

)

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Print("Некорректный ввод")
	} else if n < 10 {
		fmt.Print("Число меньше 10")
	} else if n < 100 {
		fmt.Print("Число меньше 100")
	} else if n < 1000{
		fmt.Print("Число меньше 1000")
	} else {
		fmt.Print("Число больше или равно 1000")
	}
}
