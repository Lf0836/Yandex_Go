package main

func Mix(nums []int) []int {
	n := len(nums) / 2
	result := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}
