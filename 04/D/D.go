package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	sum = 1
	for i := n; i > 0; i-- {
		sum *= i
	}
	fmt.Println(sum)
}
