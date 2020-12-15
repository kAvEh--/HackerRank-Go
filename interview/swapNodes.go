package main

import "fmt"

func main() {
	//data := [][]int32{{2, 3}, {-1, 4}, {-1, 5}, {-1, -1}, {-1, -1}}
	data := [][]int32{{2, 3}, {-1, -1}, {-1, -1}}
	queries := []int32{1, 1}
	tt := swapNodes(data, queries)
	for i := 0; i < len(tt); i++ {
		fmt.Println(tt[i])
	}
}

type Node struct {
	value int32
	depth int
	left  *Node
	right *Node
}

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	var ret [][]int32
	//var levels []int
	root := Node{
		value: 1,
		depth: 1,
	}
	c := make(chan *Node, len(indexes))
	c <- &root
	maxDepth := 0
	for i := 0; i < len(indexes); i++ {
		parent := <-c
		if maxDepth < parent.depth {
			maxDepth = parent.depth
		}
		parent.left = getNode(indexes, i, 0, parent.depth)
		parent.right = getNode(indexes, i, 1, parent.depth)

		if parent.left.value != -1 {
			c <- parent.left
		}
		if parent.right.value != -1 {
			c <- parent.right
		}
	}

	for i := 0; i < len(queries); i++ {
		loop := queries[i]
		for {
			if int(loop) >= maxDepth {
				break
			}
			swap(&root, int(loop))
			loop += queries[i]
		}

		var tmp []int32
		readTree(&root, &tmp)
		ret = append(ret, tmp)
	}
	return ret
}

func getNode(indexes [][]int32, i int, j int, depth int) *Node {
	ret := Node{
		value: indexes[i][j],
		depth: depth + 1,
	}
	return &ret
}

func readTree(n *Node, list *[]int32) {
	if n.value == -1 {
		return
	}
	if n.left != nil {
		readTree(n.left, list)
	}
	*list = append(*list, n.value)
	if n.right != nil {
		readTree(n.right, list)
	}
}

func swap(parent *Node, depth int) {
	if parent.value == -1 {
		return
	}
	if parent.depth == depth {
		parent.left, parent.right = parent.right, parent.left
		return
	}
	swap(parent.left, depth)
	swap(parent.right, depth)
}
