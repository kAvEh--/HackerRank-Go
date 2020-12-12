package main

import (
	"fmt"
)

func main() {
	data := []int32{10, 20, 30, 40, 80}
	//data := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	fmt.Println(activityNotifications(data, 2))
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	tmp := make([]int, 201)
	num := int32(0)
	for i := int(d); i < len(expenditure); i++ {
		mean := calcMean(expenditure, tmp, i, int(d))
		if float64(expenditure[i]) >= 2*mean {
			num++
		}
	}

	return num
}

func calcMean(e []int32, tmp []int, idx int, d int) float64 {
	if idx == d {
		for i := 0; i < d; i++ {
			tmp[e[i]]++
		}
	} else {
		tmp[e[idx-1]]++
		tmp[e[idx-d-1]]--
	}

	if d%2 == 1 {
		sum := 0
		for i := 0; i < len(tmp); i++ {
			sum += tmp[i]
			if sum > d/2 {
				return float64(i)
			}
		}
	} else {
		s1 := -1
		sum := 0
		for i := 0; i < len(tmp); i++ {
			sum += tmp[i]
			if s1 < 0 && sum > d/2-1 {
				s1 = i
			}
			if sum > d/2 {
				return float64(s1+i) / float64(2)
			}
		}
	}

	return 0
}
