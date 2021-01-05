package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int32{1, 3, 5, 7}
	queries := []int32{17, 6}
	ret := maxXor(data, queries)
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}

// Complete the maxXor function below.
func maxXor(arr []int32, queries []int32) []int32 {
	// solve here
	max := make([]int32, len(queries))
	m := make(map[int32]bool)
	big := int32(0)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
		if arr[i] > big {
			big = arr[i]
		}
	}
	for i := 0; i < len(queries); i++ {
		limit := int32(0)
		if queries[i] > big {
			limit = calcLimit(queries[i])
		} else {
			limit = calcLimit(big)
		}
		for j := limit; j > 0; j-- {
			if m[queries[i]^j] {
				max[i] = j
				break
			}
		}
	}
	return max
}

func calcLimit(n int32) int32 {
	tmp := int(math.Log2(float64(n))) + 1
	limit := int32(math.Pow(2, float64(tmp)))
	return limit
}
