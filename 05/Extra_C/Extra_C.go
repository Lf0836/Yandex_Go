package main

import "fmt"

func main() {
	var n, summ int
	fmt.Scan(&n)
	summ = 0
	for i := 1; i < n+1; i++ {
		if i%3 == 0 || i%5 == 0 {
		} else {
			summ += i
		}
	}
	fmt.Println(summ)
}
