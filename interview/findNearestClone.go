package main

import "fmt"

func main() {
	from := []int32{1, 1, 2, 3}
	to := []int32{2, 3, 4, 5}
	ids := []int64{1, 2, 3, 3, 2}
	fmt.Println(findShortest(5, from, to, ids, 2))
}

/*
 * For the unweighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] to <name>To[i].
 *
 */

func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64, val int32) int32 {
	// solve here
	adj := make(map[int32][]int32)
	tmp := make(map[int32][]int32)
	seen := make(map[int32][]bool)
	for i := 0; i < len(graphTo); i++ {
		adj[graphFrom[i]] = append(adj[graphFrom[i]], graphTo[i])
		adj[graphTo[i]] = append(adj[graphTo[i]], graphFrom[i])
	}
	for i := int32(0); i < graphNodes; i++ {
		if ids[i] == int64(val) {
			tmp[i+1] = make([]int32, 0)
			tmp[i+1] = append(tmp[i+1], i+1)
			seen[i+1] = make([]bool, graphNodes+1)
			seen[i+1][i+1] = true
		}
	}
	degree := int32(1)
	for len(tmp) > 0 {
		fmt.Println(">>>", seen, "::", tmp)
		t1 := make(map[int32][]int32)
		for r, k := range tmp {
			for _, tt := range k {
				for i := 0; i < len(adj[tt]); i++ {
					if !seen[r][adj[tt][i]] {
						if ids[adj[tt][i]-1] == int64(val) {
							return degree
						}
						seen[r][adj[tt][i]] = true
						t1[r] = append(t1[r], adj[tt][i])
					}
				}
			}
		}
		tmp = t1
		degree++
	}
	return -1
}
