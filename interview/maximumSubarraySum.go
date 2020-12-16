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
	mm, pr := int64(0), int64(0)
	a1 := make([]int64, 0)
	for _, i := range a {
		pr = (pr + i) % m
		mm = int64(math.Max(float64(mm), float64(pr)))
		idx := sort.Search(len(a1), func(ii int) bool { return a1[ii] >= pr+1 })
		if idx < len(a1) {
			mm = int64(math.Max(float64(mm), float64(pr-a1[idx]+m)))

			a1 = append(a1[:idx+1], a1[idx:]...)
			a1[idx] = pr
		} else {
			a1 = append(a1, pr)
		}
	}

	return mm
}
