package main

import (
	"math"
	"slices"
)

// using two pointers
func lc_squares_of_sorted_array_linear(nums []int) []int {
	result := make([]int, 0, len(nums))

	for left, right := 0, len(nums)-1; left <= right; {
		absLeft, absRight := math.Abs(float64(nums[left])), math.Abs(float64(nums[right]))

		if absLeft > absRight {
			result = append(result, int(math.Pow(absLeft, 2.0)))
			left++
		} else {
			result = append(result, int(math.Pow(absRight, 2.0)))
			right--
		}
	}

	slices.Reverse(result)

	return result
}

// squaring and sorting after, squaring is linear, sorting is "n * log(n)"
func lc_squares_of_sorted_array_nlogn(nums []int) []int {
	result := make([]int, 0, len(nums))

	for n := range nums {
		result = append(result, int(math.Pow(float64(nums[n]), 2)))
	}

	slices.Sort(result)

	return result
}
