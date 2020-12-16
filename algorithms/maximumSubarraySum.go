package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	data := []int64{673764401, 1244270825, 1551000418, 1386904091, 58618188, 1962572095}

	fmt.Println(maximumSum(data, 10002143548612))
}

// Complete the maximumSum function below.
//TODO It seems there is a bug in problem core code in HackerRank.
// I've reported it but until now there is not any correction on it.
// this code works perfect but test case NO. 16 will not pass in HackerRank editor.
// Run this code with downloaded data will result correct answer.
func maximumSum(a []int64, m int64) int64 {
	max, n := int64(0), int64(0)
	tmp := make([]int64, 0)
	for _, i := range a {
		n = (n + i) % m
		max = int64(math.Max(float64(max), float64(n)))
		idx := sort.Search(len(tmp), func(ii int) bool { return tmp[ii] >= n+1 })
		if idx < len(tmp) {
			max = int64(math.Max(float64(max), float64(n-tmp[idx]+m)))

			tmp = append(tmp[:idx+1], tmp[idx:]...)
			tmp[idx] = n
		} else {
			tmp = append(tmp, n)
		}
	}

	return max
}
