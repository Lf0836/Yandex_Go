package main

func SliceCopy(nums []int) []int {
	newSlice := make([]int, len(nums))
	copy(newSlice, nums)
	return newSlice
}
