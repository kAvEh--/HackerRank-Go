package main

import (
	"fmt"
	"sort"
)

func main() {
	data := [][]int32{{5, 1}, {1, 1}, {4, 0}}
	fmt.Println(luckBalance(2, data))
}

// Complete the luckBalance function below.
func luckBalance(k int32, contests [][]int32) int32 {
	sum := int32(0)

	sort.SliceStable(contests, func(i, j int) bool {
		if contests[i][1] == 0 {
			return false
		}
		if contests[j][1] == 0 {
			return true
		}
		return contests[i][0] < contests[j][0]
	})

	count := 0
	for i := 0; i < len(contests); i++ {
		sum += contests[i][0]
		if contests[i][1] == 1 {
			count++
		}
	}

	for i := 0; i < count-int(k); i++ {
		sum -= 2 * contests[i][0]
	}

	return sum
}
