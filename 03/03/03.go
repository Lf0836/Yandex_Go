package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Print("Число положи")
	} else {
		fmt.Print("Число отрица")
	}
	if n == 0 {
		fmt.Print("Число равно нулю")
	}
	if n%2 == 0 {
		fmt.Print("тельное и четное")
	} else {
		fmt.Print("тельное и нечетное")
	}
}
