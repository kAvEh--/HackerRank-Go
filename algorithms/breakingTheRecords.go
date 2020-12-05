package main

import "fmt"

func main() {
	data := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	ret := breakingRecords(data)
	fmt.Println(ret)
}

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	var ret []int32
	ret = append(ret, 0, 0)
	min, max := scores[0], scores[0]
	for i := 1; i < len(scores); i++ {
		if min > scores[i] {
			min = scores[i]
			ret[1]++
		}
		if max < scores[i] {
			max = scores[i]
			ret[0]++
		}
	}
	return ret
}
