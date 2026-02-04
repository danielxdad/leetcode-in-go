package main

import (
	"fmt"
)

// https://leetcode.com/problems/container-with-most-water/
func leetcode_11(height []int) int {
	// The maximum amount of water that can hold a bucket delimited by (height[x], height[y])
	// is min(height[x], height[y]) times diff between offset i and j (j - i)
	// area = min(height[i], height[j]) * (j - i)
	minimun := min(height[0], height[len(height)-1])
	max_area := minimun * (len(height) - 1)
	max_i, max_j := 0, len(height)-1

	// Two pointers(fingers)
	for i, j := 0, len(height)-1; i < j; i, j = i+1, j-1 {
		for n_i := i + 1; n_i < j; n_i++ {
			if height[n_i] >= (height[i] + (n_i - i)) {
				new_area := min(height[n_i], height[j]) * (j - n_i)
				if new_area > max_area {
					i = n_i
					max_area = new_area
					max_i = i
				}
			}
		}
		for n_j := j - 1; n_j > i; n_j-- {
			if height[n_j] >= (height[j] + (j - n_j)) {
				new_area := min(height[i], height[n_j]) * (n_j - i)
				if new_area > max_area {
					j = n_j
					max_area = new_area
					max_j = j
				}
			}
		}
	}

	// i, j := 0, len(height)-1
	// max_area := min(height[i], height[j]) * j
	//
	// for ; i < len(height) && j > -1; i, j = i+1, j-1 {
	// 	lower, upper := min(i, j), max(i, j)
	// 	offset := upper - lower
	// 	lower_ptr := &i
	// 	upper_ptr := &j
	//
	// 	if i > j {
	// 		lower_ptr = &j
	// 		upper_ptr = &i
	// 	}
	//
	// 	for n_i := lower; n_i < upper; n_i++ {
	// 		if height[n_i] > (height[*lower_ptr] + (n_i - lower)) {
	// 			*lower_ptr = n_i
	// 			max_area = max(max_area, min(height[*lower_ptr], height[*upper_ptr])*offset)
	// 		}
	// 	}
	//
	// 	for n_j := upper; n_j > lower; n_j-- {
	// 		if height[n_j] > (height[*upper_ptr] + (*upper_ptr - n_j)) {
	// 			*upper_ptr = n_j
	// 			max_area = max(max_area, min(height[*lower_ptr], height[*upper_ptr])*offset)
	// 		}
	// 	}
	// }
	//
	fmt.Println(height)
	fmt.Println(max_area, max_i, max_j, height[max_i], height[max_j])

	return max_area
}

func leetcode_11_1(height []int) int {
	// Bruteforce
	max_area := 0
	max_i := -1
	max_j := -1
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			// max_area = max(max_area, min(height[i], height[j])*(j-i))
			na := min(height[i], height[j]) * (j - i)
			if na > max_area {
				max_area = na
				max_i = i
				max_j = j
			}
		}
	}

	// for i, j := 0, 1; i != j; i, j = (i+1)%len(height), (j+2)%len(height) {
	// 	offset := max(j, i) - min(j, i)
	// 	max_area = max(max_area, min(height[i], height[j])*offset)
	// 	fmt.Println(i, j, offset, height[i], height[j])
	// }

	fmt.Println("BF:", height)
	fmt.Println("BF:", max_area, max_i, max_j, height[max_i], height[max_j])

	return max_area
}
