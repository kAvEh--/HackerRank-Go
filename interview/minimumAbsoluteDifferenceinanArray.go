package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	data := []int32{1, -3, 71, 68, 17}
	fmt.Println(minimumAbsoluteDifference(data))
}

// Complete the minimumAbsoluteDifference function below.
func minimumAbsoluteDifference(arr []int32) int32 {
	var min = int32(math.MaxInt32)
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1] - arr[i]
		}
	}
	return min
}
