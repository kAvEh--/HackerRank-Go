package main

import "fmt"

func main() {
	fmt.Println(gridSearch([]string{
		"7283455864",
		"6731158619",
		"8588242643",
		"3830589324",
		"2229505813",
		"5633845374",
		"6473530293",
		"7053106601",
		"0834282956",
		"4607924137",
	}, []string{
		"9505",
		"3845",
		"3530",
	}))
}

/*
 * Complete the 'gridSearch' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY G
 *  2. STRING_ARRAY P
 */

func gridSearch(G []string, P []string) string {
	for i := 0; i < len(G)-len(P); i++ {
		for j := 0; j < len(G[0])-len(P[0]); j++ {
			if G[i][j] == P[0][0] {
				if check(i, j, G, P) {
					return "YES"
				}
			}
		}
	}

	return "NO"
}

func check(i1 int, j1 int, g []string, p []string) bool {
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			if g[i1+i][j1+j] != p[i][j] {
				return false
			}
		}
	}
	return true
}
