package main

func Clean(nums []int, x int) []int {
	index := 0
	for _, num := range nums {
		if num != x {
			nums[index] = num
			index++
		}
	}
	return nums[:index]
}
