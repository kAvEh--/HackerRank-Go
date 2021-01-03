package main

import (
	"fmt"
)

func main() {
	data := []int32{4, 5, 4, 6, 4, 5, 1}
	fmt.Println(candies(10, data))
}

// Complete the candies function below.
func candies(n int32, arr []int32) int64 {
	cand := make([]int64, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			cand[i+1] = cand[i] + 1
		}
	}
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] && cand[i-1] <= cand[i] {
			cand[i-1] = cand[i] + 1
		}
	}
	sum := int64(0)
	for i := 0; i < len(cand); i++ {
		sum++
		sum += cand[i]
	}
	return sum
}
