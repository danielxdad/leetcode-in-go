package main

import (
	"fmt"
	"maps"
	// "slices"
)

func main() {
	// leetcode_217()

	// =========================================================
	// fmt.Println(leetcode_242("anagram", "nagaram"))
	// fmt.Println(leetcode_242("rat", "car"))
	// fmt.Println(leetcode_242("aacc", "ccac"))

	// =========================================================
	// height := []int{1, 1}
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{4, 3, 2, 1, 4}
	// height := []int{1, 2, 1}
	// height := []int{8, 7, 2, 1}
	// height := []int{1, 2, 3, 4, 5, 25, 24, 3, 4}
	// height := []int{5, 2, 12, 1, 5, 3, 4, 11, 9, 4}
	// height := []int{1, 2, 4, 3}
	// height := []int{4, 6, 4, 4, 6, 2, 6, 7, 11, 2}
	// fmt.Println(leetcode_11(height))
	// fmt.Println(leetcode_11_1(height))

	// =========================================================
	// fmt.Println(zigzag("PAYPALISHIRING", 3))
	// fmt.Println(zigzag("PAYPALISHIRING", 4))
	// fmt.Println(zigzag("A", 1))
	// fmt.Println(zigzag("A", 2))
	// fmt.Println(zigzag("ABCD", 3))
	// fmt.Println(zigzag("ABCD", 2))

	// =========================================================
	// data := make([]int, 0, 1000)
	// // data := []int{-4, -1, 0, 3, 10}

	// for range cap(data) {
	// 	if rand.Intn(2) > 0 {
	// 		data = append(data, rand.Intn(cap(data))*-1)
	// 	} else {
	// 		data = append(data, rand.Intn(cap(data)))
	// 	}
	// }

	// slices.Sort(data)
	// // data = slices.Compact(data)
	// // fmt.Println(len(data), cap(data))
	// fmt.Println(data[0:5], "...", data[len(data)-5:])

	// n := 100_000
	// start := time.Now()
	// for range n {
	// 	lc_squares_of_sorted_array_linear(data)
	// }
	// linearDuration := time.Since(start)
	// fmt.Println(linearDuration)

	// start = time.Now()
	// for range n {
	// 	lc_squares_of_sorted_array_nlogn(data)
	// }
	// nlognDuration := time.Since(start)
	// fmt.Println(nlognDuration)

	// fmt.Println("nlogn / linear: ", float64(nlognDuration.Milliseconds())/float64(linearDuration.Milliseconds()))
	// ===============================================================================================================

	fmt.Println(distribute_elements_into_two_arrays_3069([]int{2, 1, 3}))
	fmt.Println(distribute_elements_into_two_arrays_3069([]int{5, 4, 3, 8}))
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
