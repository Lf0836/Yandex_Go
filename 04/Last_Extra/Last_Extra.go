package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n)
	for b = 1; b < n; a, b = b, a+b {
	}
	for i := 0; i < 10; i++ {
		a, b = b, a+b
		fmt.Println(a)
	}
}
