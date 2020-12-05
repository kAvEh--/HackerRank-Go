package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	data := []int32{1, 5, 4, 3, 2, 6}
	almostSorted(data)
}

// Complete the almostSorted function below.
func almostSorted(arr []int32) {
	sorted := make([]int32, len(arr))
	copy(sorted, arr)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	first, last := -1, -1
	for i := 0; i < len(arr); i++ {
		if arr[i] != sorted[i] {
			if first == -1 {
				first = i
			} else {
				last = i
			}
		}
	}
	if first == -1 && last == -1 {
		fmt.Println("yes")
		return
	}
	arr[first], arr[last] = arr[last], arr[first]
	if reflect.DeepEqual(sorted, arr) {
		fmt.Println("yes")
		fmt.Println("swap", first+1, last+1)
		return
	}
	arr[first], arr[last] = arr[last], arr[first]
	for i := 0; i < (last-first+1)/2; i++ {
		arr[first+i], arr[last-i] = arr[last-i], arr[first+i]
	}
	//fmt.Println(arr)
	//fmt.Println(sorted)
	if reflect.DeepEqual(sorted, arr) {
		fmt.Println("yes")
		fmt.Println("reverse", first+1, last+1)
		return
	}
	fmt.Println("no")
}
