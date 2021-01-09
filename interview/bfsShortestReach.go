package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, _ := strconv.ParseInt(readLine1(reader), 10, 64)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine1(reader)

		tmp := strings.Fields(s)
		nodes, _ := strconv.Atoi(tmp[0])
		edge, _ := strconv.Atoi(tmp[1])
		adj := make([][]int, nodes)
		for i := range adj {
			adj[i] = make([]int, nodes)

		}
		for i := 0; i < edge; i++ {
			s = readLine1(reader)
			tmp = strings.Fields(s)
			t1, _ := strconv.Atoi(tmp[0])
			t2, _ := strconv.Atoi(tmp[1])
			adj[t1-1][t2-1] = 1
			adj[t2-1][t1-1] = 1
		}
		s = readLine1(reader)
		tt, _ := strconv.Atoi(s)
		bfs(nodes, adj, tt-1)
	}
}

func bfs(node int, adj [][]int, start int) {
	nodes := make([]int, node)
	tmp := make(map[int]bool)
	tmp[start] = true
	distance := 1
	for len(tmp) > 0 {
		t1 := make(map[int]bool)
		for i := range tmp {
			for j := 0; j < node; j++ {
				if adj[i][j] == 1 && nodes[j] < 1 {
					t1[j] = true
					nodes[j] = distance * 6
				}
			}
		}
		tmp = t1
		distance++
	}
	for i := 0; i < node; i++ {
		if i != start {
			if nodes[i] == 0 {
				fmt.Print("-1 ")
			} else {
				fmt.Print(nodes[i], " ")
			}
		}
	}
	fmt.Println()
}

func readLine1(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
