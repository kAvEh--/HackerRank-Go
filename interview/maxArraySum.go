package main

import "fmt"

func main() {
	data := []int32{3, 7, 4, 6, 5}
	fmt.Println(maxSubsetSum(data))
}

var memo map[int]int32

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	memo = make(map[int]int32)
	sum := max(arr)
	return sum
}

func max(arr []int32) int32 {
	if len(arr) < 1 {
		return 0
	}
	if len(arr) == 1 {
		if arr[0] > 0 {
			return arr[0]
		}
		return 0
	}
	tmp, ok := memo[len(arr)]
	if ok {
		return tmp
	}
	a, b := int32(0), int32(0)
	tmp, ok = memo[len(arr)-1]
	if ok {
		a = tmp
	} else {
		a = max(arr[1:])
		memo[len(arr)-1] = a
	}
	memo[len(arr)] = a
	if arr[0] > 0 {
		tmp, ok := memo[len(arr)-2]
		if ok {
			b = tmp
		} else {
			b = max(arr[2:])
			memo[len(arr)-2] = b
		}
		if a >= b+arr[0] {
			return a
		}
		memo[len(arr)] = b + arr[0]
		return b + arr[0]
	}
	return a
}
