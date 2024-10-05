package main

func FiveSteps(array [5]int) [5]int {
	var reversed [5]int
	for i := 0; i < 5; i++ {
		reversed[i] = array[4-i]
	}

	return reversed
}
