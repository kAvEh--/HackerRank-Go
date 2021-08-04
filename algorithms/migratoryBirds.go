package main

import "fmt"

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	birds := make(map[int32]int)
	for i := 0; i < len(arr); i++ {
		birds[arr[i]]++
	}
	km, max := int32(-1), 0
	for k, v := range birds {
		if km == -1 {
			km = k
			max = v
			continue
		}
		if v > max {
			max = v
			km = k
			continue
		}
		if v == max && km > k {
			km = k
		}
	}
	return km
}
