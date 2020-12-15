package main

import "fmt"

func main() {
	data := []int64{3, 3, 9, 9, 5}
	fmt.Println(maximumSum(data, 7))
}

// Complete the maximumSum function below.
//TODO optimization needed current order is O(n^2)
// 		must reduce to O(nlgn)
func maximumSum(a []int64, m int64) int64 {
	max := int64(0)
	modulo := make(map[int64]int)
	for i := 0; i < len(a); i++ {
		tmp := a[i] % m
		modulo[tmp] = 1
		if tmp > max {
			if tmp == m-1 {
				return tmp
			}
			max = tmp
		}
		for j := tmp + 1; j < m; j++ {
			if (tmp-j+m)%m < max {
				break
			}
			if _, ok := modulo[j]; ok {
				max = (tmp - j + m) % m
				break
			}
		}
		if tmp == m-1 {
			return tmp
		}
	}

	return max
}
