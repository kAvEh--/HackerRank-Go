package main

import "fmt"

func main() {
	fmt.Println(absolutePermutation(95, 22))
}

func absolutePermutation(n int32, k int32) []int32 {
	// Write your code here
	if k == 0 {
		ret := make([]int32, n)
		for i := int32(0); i < n; i++ {
			ret[i] = i + 1
		}
		return ret
	}
	if n%(2*k) != 0 || n/2 < k {
		ret := make([]int32, 1)
		ret[0] = -1
		return ret
	}
	ret := make([]int32, n)
	for i := int32(0); i < n; i++ {
		ret[i] = i + 1
	}
	ct := make([]bool, n)
	for i := int32(0); i < n-k; i++ {
		if !ct[i] {
			ret[i], ret[i+k] = ret[i+k], ret[i]
			ct[i+k] = true
		}
	}

	return ret
}
