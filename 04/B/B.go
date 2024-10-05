package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < 11; i++ {
		fmt.Print(n)
		fmt.Print(" * ")
		fmt.Print(i)
		fmt.Print(" = ")
		fmt.Print(n * i)
	}
}
