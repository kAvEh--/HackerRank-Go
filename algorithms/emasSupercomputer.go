package main

import (
	"fmt"
)

func main() {
	data := []string{"GGGGGGGGGGGG", "GBGGBBBBBBBG", "GBGGBBBBBBBG", "GGGGGGGGGGGG",
		"GGGGGGGGGGGG", "GGGGGGGGGGGG", "GGGGGGGGGGGG", "GBGGBBBBBBBG", "GBGGBBBBBBBG",
		"GBGGBBBBBBBG", "GGGGGGGGGGGG", "GBGGBBBBBBBG"}
	fmt.Println(twoPluses(data))
}

// Complete the twoPluses function below.
func twoPluses(grid []string) int32 {
	table := make([][]int, len(grid))
	for i := range table {
		table[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "B" {
				table[i][j] = 1
			}
		}
	}
	min := len(table)
	if min > len(table[0]) {
		min = len(table[0])
	}
	if min%2 == 0 {
		min--
	}
	var plus [][]int
	for i := min; i > 0; i = i - 2 {
		for j := 0; j <= len(table)-i; j++ {
			for k := 0; k <= len(table[0])-i; k++ {
				if checkPlus(table, i, j, k) {
					plus = append(plus, []int{i, j, k})
				}
			}

		}
	}

	max := 0
	for i := 0; i < len(plus); i++ {
		for j := i + 1; j < len(plus); j++ {
			if (2*plus[i][0]-1)*(2*plus[j][0]-1) > max {
				if !checkOverlap(table, plus[i], plus[j]) {
					max = (2*plus[i][0] - 1) * (2*plus[j][0] - 1)
				}
			}
		}
	}

	return int32(max)
}

func checkOverlap(table [][]int, first []int, second []int) bool {
	tmp := make([][]int, len(table))
	for i := 0; i < len(table); i++ {
		tt := make([]int, len(table[i]))
		for j := 0; j < len(table[i]); j++ {
			tt[j] = table[i][j]
		}
		tmp[i] = tt
	}
	for r := 0; r < first[0]; r++ {
		tmp[first[1]+r][first[2]+(first[0]/2)] = 2
	}
	for r := 0; r < first[0]; r++ {
		tmp[first[1]+(first[0]/2)][first[2]+r] = 2
	}
	for i := 0; i < second[0]; i++ {
		if tmp[second[1]+i][second[2]+(second[0]/2)] == 2 {
			return true
		}
		tmp[second[1]+i][second[2]+(second[0]/2)] = 3
	}
	for i := 0; i < second[0]; i++ {
		if tmp[second[1]+(second[0]/2)][second[2]+i] == 2 {
			return true
		}
		tmp[second[1]+(second[0]/2)][second[2]+i] = 3
	}
	return false
}

func checkPlus(table [][]int, i int, j int, k int) bool {
	for r := 0; r < i; r++ {
		if table[j+(i/2)][k+r] != 0 {
			return false
		}
	}
	for r := 0; r < i; r++ {
		if table[j+r][k+(i/2)] != 0 {
			return false
		}
	}
	return true
}
