package main

import "fmt"

func main() {
	data := []int32{4}
	ret := birthday(data, 4, 1)
	fmt.Println(ret)
}

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {
	//var sum int32 = 0
	sum := make([]int32, len(s))
	sum[0] = s[0]
	for i := 1; i < len(s); i++ {
		sum[i] = sum[i-1] + s[i]
	}
	var ret int32 = 0
	if sum[int(m)-1] == d {
		ret++
	}
	for i := 0; i < len(s)-int(m); i++ {
		if sum[i+int(m)]-sum[i] == d {
			ret++
		}
	}

	return ret
}
