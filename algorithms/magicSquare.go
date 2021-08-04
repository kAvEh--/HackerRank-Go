package main

import (
	"fmt"
)

func main() {
	fmt.Println(formingMagicSquare([][]int32{
		{4, 5, 8},
		{2, 4, 1},
		{1, 9, 7},
	}))
}

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func formingMagicSquare(s [][]int32) int32 {
	all := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}}, // 1

		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}}, // 2

		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}}, // 3

		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}}, // 4

		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}}, // 5

		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}}, // 6

		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}}, // 7

		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}}, // 8
	}

	min := int32(1000000)
	for i := 0; i < len(all); i++ {
		tmp := int32(0)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				tmp += diff(all[i][j][k], s[j][k])
			}
		}
		if tmp < min {
			min = tmp
		}
	}
	return min
}

func diff(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}
