package main

import "fmt"

func main() {
	data := []int32{1, 3, 2, 6, 1, 2}
	ret := divisibleSumPairs(6, 3, data)
	fmt.Println(ret)
}

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	ret := int32(0)
	for i := 1; i < int(n); i++ {
		for j := 0; j < i; j++ {
			if (ar[i]+ar[j])%k == 0 {
				ret++
			}
		}
	}
	return ret
}
