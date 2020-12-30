package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	data := []int32{10, 100, 300, 200, 1000, 20, 30}
	fmt.Println(maxMin(3, data))
}

// Complete the maxMin function below.
func maxMin(k int32, arr []int32) int32 {
	min := int32(math.MaxInt32)
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	for i := 0; i <= len(arr)-int(k); i++ {
		if arr[i+int(k)-1]-arr[i] < min {
			min = arr[i+int(k)-1] - arr[i]
		}
	}

	return min
}
