package main

// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
func distribute_elements_into_two_arrays_3069(arr []int) []int {
	var arr1 []int
	var arr2 []int

	arr1 = append(arr1, arr[0])
	arr2 = append(arr2, arr[1])

	for i := 2; i < len(arr); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, arr[i])
		} else {
			arr2 = append(arr2, arr[i])
		}
	}

	return append(arr1, arr2...)
}
