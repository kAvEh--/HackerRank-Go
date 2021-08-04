package main

import "fmt"

func main() {
	fmt.Println(climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}))
}

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	ret := make([]int32, 0)
	tmp := make([]int32, 0)
	tmp = append(tmp, ranked[0])
	for i := 1; i < len(ranked); i++ {
		if ranked[i] != ranked[i-1] {
			tmp = append(tmp, ranked[i])
		}
	}

	i := int32(len(tmp) - 1)
	for _, score := range player {
		for i >= 0 {
			if score >= tmp[i] {
				i--
			} else {
				ret = append(ret, i+2)
				break
			}
		}
		if i < 0 {
			ret = append(ret, 1)
		}
	}

	return ret
}
