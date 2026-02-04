package main

import (
	"slices"
)

// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/?envType=daily-question&envId=2026-01-22
func minimumPairRemoval(nums []int) int {
	c := 0

	for len(nums) > 1 {
		lsi := 0
		least_sum := nums[lsi] + nums[lsi+1]
		non_dec := nums[lsi] <= nums[lsi+1]

		for i := 1; i < len(nums)-1; i++ {
			non_dec = non_dec && (nums[i] <= nums[i+1])
			if least_sum > nums[i]+nums[i+1] {
				least_sum = nums[i] + nums[i+1]
				lsi = i
			}
		}

		if non_dec {
			break
		}

		nums[lsi] = nums[lsi] + nums[lsi+1]
		nums = slices.Delete(nums, lsi+1, lsi+2)
		c++
	}

	return c
}
