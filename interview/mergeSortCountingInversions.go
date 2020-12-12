package main

import (
	"fmt"
)

func main() {
	data := []int32{2, 1, 3, 1, 2}
	fmt.Println(countInversions(data))
}

// Complete the countInversions function below.
func countInversions(arr []int32) int64 {
	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int32, left int, right int) int64 {
	count := int64(0)
	if left < right {
		mid := (left + right) / 2

		count += mergeSort(arr, left, mid)

		count += mergeSort(arr, mid+1, right)

		count += mergeCount(arr, left, mid, right)
	}

	return count
}

func mergeCount(arr []int32, left int, mid int, right int) int64 {
	tmp := make([]int32, right-left+1)

	i, j, k, swaps := 0, 0, left, int64(0)
	for i < mid-left+1 && j < right-mid {
		if arr[left+i] <= arr[mid+1+j] {
			tmp[k-left] = arr[left+i]
			i++
		} else {
			tmp[k-left] = arr[mid+1+j]
			j++
			swaps += int64((mid + 1) - (left + i))
		}
		k++
	}
	for i < mid-left+1 {
		tmp[k-left] = arr[left+i]
		k++
		i++
	}
	for j < right-mid {
		tmp[k-left] = arr[mid+1+j]
		k++
		j++
	}
	for m := 0; m < k-left; m++ {
		arr[left+m] = tmp[m]
	}
	return swaps
}
