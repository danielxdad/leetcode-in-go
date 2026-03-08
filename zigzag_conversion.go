package main

import (
	// "fmt"
	"math"
)

// https://leetcode.com/problems/zigzag-conversion/
func zigzag(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// result := ""
	// rows := make(map[int]string, numRows)
	// isCol := true
	// for i := 0; i < len(s); i += numRows - 1 {
	// 	batch := s[i:min(i+numRows, len(s))]
	//
	// 	for j, c := range batch {
	// 		if isCol {
	// 			rows[j] += string(c)
	// 		} else {
	// 			if j == 0 || (j == len(batch)-1 && len(batch) == numRows) {
	// 				continue
	// 			}
	// 			rows[numRows-j-1] += string(c)
	// 		}
	// 	}
	//
	// 	isCol = !isCol
	// }
	//
	// for i := range numRows {
	// 	result += rows[i]
	// }
	// ===========================================================================

	// abs(((n - 1) - (10 % (n - 1))) * 2)
	// abs(((numRows - 1) - (index % (numRows - 1)))) * 2
	result := make([]byte, 0, len(s))
	for i := range numRows {
		for j := i; j < len(s); {
			result = append(result, s[j])

			// numRows - 1: The last index of the first chunk/column of length "numRows"
			// j % (numRows - 1): scaled the "j" index into the range [0..(numRows - 1))
			// that give you the number of elements to the "corner index" or "last index"
			// of a chunk of length "numRows" and times 2 give the next index
			// where we can find the next character for this "row"
			j += int(math.Abs(float64(numRows-1)-float64(j%(numRows-1)))) * 2
		}
	}

	return string(result)
}

/*
P   A   H   N
A P L S I I G
Y   I   R

0, 4, 8, 12		: 4
1, 3, 5, 7, 9, 11, 13	: 2
2, 6, 10		: 4
*/

/*
P     I     N
A   L S   I G
Y A   H R
P     I

0, 6, 12 		: 6
1, 5, 7, 11, 13		: 4, 2
2, 4, 8, 10		: 2, 4
3, 9			: 6
*/

/*
"ABCD", 3

A
B D
C
*/
