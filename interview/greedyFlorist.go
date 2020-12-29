package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int32{2, 5, 6}
	fmt.Println(getMinimumCost(2, data))
}

// Complete the getMinimumCost function below.
func getMinimumCost(k int32, c []int32) int32 {
	cost := int32(0)
	sort.SliceStable(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	round := int32(0)
	for int(round*k) < len(c) {
		for i := 0; i < int(k); i++ {
			tmp := int(round*k) + i
			if tmp >= len(c) {
				break
			}
			cost += (round + 1) * c[tmp]
		}
		round++
	}
	return cost
}
