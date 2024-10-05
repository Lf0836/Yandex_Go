package main

import (
	"fmt"
	"math/rand"
	"time"

)

func FindMinMaxInArray(array [10]int) (int, int) {
	min := array[0]
	max := array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100) 
	}
	fmt.Println("Сгенерированный массив:", array)
	min, max := FindMinMaxInArray(array)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Максимальное значение: %d\n", max)
}