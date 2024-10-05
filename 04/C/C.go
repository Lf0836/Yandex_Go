package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		sum += i
	}
	fmt.Println(sum)
}
