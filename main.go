package main

import (
	"fmt"
	"maps"
	// "slices"
)

func main() {
	// leetcode_217()

	// fmt.Println(leetcode_242("anagram", "nagaram"))
	// fmt.Println(leetcode_242("rat", "car"))
	// fmt.Println(leetcode_242("aacc", "ccac"))

	// height := []int{1, 1}
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{4, 3, 2, 1, 4}
	// height := []int{1, 2, 1}
	// height := []int{8, 7, 2, 1}
	height := []int{1, 2, 3, 4, 5, 25, 24, 3, 4}
	// height := []int{5, 2, 12, 1, 5, 3, 4, 11, 9, 4}
	// height := []int{1, 2, 4, 3}
	// height := []int{4, 6, 4, 4, 6, 2, 6, 7, 11, 2}
	fmt.Println(leetcode_11(height))
	fmt.Println(leetcode_11_1(height))
}

func leetcode_217() bool {
	// https://leetcode.com/problems/contains-duplicate/description/
	nums := []int{1, 2, 3, 1}
	m := make(map[int]bool, len(nums))

	for i := range len(nums) {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = true
	}

	return false
}

func leetcode_242(s string, t string) bool {
	// https://leetcode.com/problems/valid-anagram/
	if len(s) != len(t) {
		return false
	}

	cs := make(map[uint8]int, len(s))
	ct := make(map[uint8]int, len(s))

	for i := range len(s) {
		cs[s[i]] += 1
		ct[t[i]] += 1
	}

	return maps.Equal(cs, ct)
}
