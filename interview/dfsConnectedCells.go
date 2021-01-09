package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := [][]int32{{0, 1, 0, 0, 0, 0, 1, 1, 0}, {1, 1, 0, 0, 1, 0, 0, 0, 1}, {0, 0, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 1, 0, 1, 1}, {0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 0, 1, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 1, 0, 1, 1}, {1, 0, 1, 1, 1, 1, 0, 0, 0}}
	fmt.Println(maxRegion(data))
}

// Complete the maxRegion function below.
func maxRegion(grid [][]int32) int32 {
	max := int32(0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				reg := clearRegion(grid, i, j)
				for i := 0; i < len(grid); i++ {
					//fmt.Println(grid[i])
				}
				//fmt.Println("----------------")
				if reg > max {
					max = reg
				}
			}
		}
	}

	return max
}

func clearRegion(grid [][]int32, ii int, jj int) int32 {
	tmp := make(map[string][]int)
	fuck(tmp, ii, jj)
	region := 1
	for len(tmp) > 0 {
		for key, node := range tmp {
			grid[node[0]][node[1]] = 2
			delete(tmp, key)
			//first row
			if isOk(node[0]-1, node[1]-1, len(grid), len(grid[0])) && grid[node[0]-1][node[1]-1] == 1 {
				grid[node[0]-1][node[1]-1] = 2
				fuck(tmp, node[0]-1, node[1]-1)
				region++
			}
			if isOk(node[0]-1, node[1], len(grid), len(grid[0])) && grid[node[0]-1][node[1]] == 1 {
				grid[node[0]-1][node[1]] = 2
				fuck(tmp, node[0]-1, node[1])
				region++
			}
			if isOk(node[0]-1, node[1]+1, len(grid), len(grid[0])) && grid[node[0]-1][node[1]+1] == 1 {
				grid[node[0]-1][node[1]+1] = 2
				fuck(tmp, node[0]-1, node[1]+1)
				region++
			}
			//second row
			if isOk(node[0], node[1]-1, len(grid), len(grid[0])) && grid[node[0]][node[1]-1] == 1 {
				grid[node[0]][node[1]-1] = 2
				fuck(tmp, node[0], node[1]-1)
				region++
			}
			if isOk(node[0], node[1]+1, len(grid), len(grid[0])) && grid[node[0]][node[1]+1] == 1 {
				grid[node[0]][node[1]+1] = 2
				fuck(tmp, node[0], node[1]+1)
				region++
			}
			// third row
			if isOk(node[0]+1, node[1]-1, len(grid), len(grid[0])) && grid[node[0]+1][node[1]-1] == 1 {
				grid[node[0]+1][node[1]-1] = 2
				fuck(tmp, node[0]+1, node[1]-1)
				region++
			}
			if isOk(node[0]+1, node[1], len(grid), len(grid[0])) && grid[node[0]+1][node[1]] == 1 {
				grid[node[0]+1][node[1]] = 2
				fuck(tmp, node[0]+1, node[1])
				region++
			}
			if isOk(node[0]+1, node[1]+1, len(grid), len(grid[0])) && grid[node[0]+1][node[1]+1] == 1 {
				grid[node[0]+1][node[1]+1] = 2
				fuck(tmp, node[0]+1, node[1]+1)
				region++
			}
		}
	}
	return int32(region)
}

func fuck(tmp map[string][]int, i int, j int) {
	tt := make([]int, 2)
	tt[0], tt[1] = i, j
	tmp[strconv.Itoa(i)+":"+strconv.Itoa(j)] = tt
}

func isOk(i int, j int, n int, m int) bool {
	return i > -1 && i < n && j > -1 && j < m
}
