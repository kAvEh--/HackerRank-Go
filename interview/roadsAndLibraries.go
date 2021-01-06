package main

import "fmt"

func main() {
	data := [][]int32{{1, 2}, {3, 1}, {2, 3}}
	fmt.Println(roadsAndLibraries(96295, 81523, 98351, data))
}

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	cost := int64(0)
	if c_road >= c_lib {
		return int64(n) * int64(c_lib)
	}
	visited = make([]int, n+1)
	adj = make(map[int32][]int32)
	for i := 0; i < len(cities); i++ {
		adj[cities[i][0]] = append(adj[cities[i][0]], cities[i][1])
		adj[cities[i][1]] = append(adj[cities[i][1]], cities[i][0])
	}
	circle := 0
	circleCount = make(map[int]int)
	for i := 1; i <= int(n); i++ {
		if visited[i] == 0 {
			circle++
			visit(circle, int32(i))
		}
	}
	cost = int64(c_lib) * int64(circle)
	for i := 0; i < circle; i++ {
		cost += int64(circleCount[i+1]-1) * int64(c_road)
	}

	return cost
}

var visited []int
var adj map[int32][]int32
var circleCount map[int]int

func visit(c int, i int32) {
	if visited[i] != 0 {
		return
	}
	visited[i] = c
	circleCount[c]++
	for _, n := range adj[i] {
		if visited[n] == 0 {
			visit(c, n)
		}
	}
}
