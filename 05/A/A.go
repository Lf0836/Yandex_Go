package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	D := b*b - 4*a*c
	if D > 0 {
		x1 := (-b - math.Sqrt(D)) / (2 * a)
		x2 := (-b + math.Sqrt(D)) / (2 * a)
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		fmt.Printf("%.4f %.4f\n", x1, x2)
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("%.4f\n", x)
	} else {
		fmt.Println("0 0")
	}
}

func main() {
	SqRoots()
}
