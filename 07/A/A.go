package main

import "errors"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n должно быть неотрицательным")
	}
	if nums == nil {
		return nil, errors.New("nums не может быть nil")
	}
	ans := make([]int, 0, n)
	for _, num := range nums {
		if num < limit {
			ans = append(ans, num)
			if len(ans) == n {
				break
			}
		}
	}
	return ans, nil
}
