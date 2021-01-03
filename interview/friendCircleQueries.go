package main

import "fmt"

func main() {
	//data := [][]int32{{6, 4}, {5, 9}, {8, 5}, {4, 1}, {1, 5}, {7, 2}, {4, 2}, {7, 6}}
	data := [][]int32{{1, 2}, {3, 4}, {1, 3}, {5, 7}, {5, 6}, {7, 4}}
	ret := maxCircle(data)
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}

// Complete the maxCircle function below.
func maxCircle(queries [][]int32) []int32 {
	m1 := make(map[int32]int)
	m2 := make(map[int]int)
	m3 := make(map[int][]int32)
	idx := 1
	max := 1
	ret := make([]int32, len(queries))
	for i := 0; i < len(queries); i++ {
		if m1[queries[i][0]] > 0 {
			if m1[queries[i][1]] > 0 {
				if m1[queries[i][0]] != m1[queries[i][1]] {
					t1, t2 := m1[queries[i][0]], m1[queries[i][1]]
					if len(m3[t1]) > len(m3[t2]) {
						for j := range m3[t2] {
							m2[t2]--
							m1[m3[t2][j]] = t1
							m2[t1]++
							m3[t1] = append(m3[t1], m3[t2][j])
						}
						m3[t2] = make([]int32, 0)
						if m2[t1] > m2[max] {
							max = t1
						}
					} else {
						for j := range m3[t1] {
							m2[t1]--
							m1[m3[t1][j]] = t2
							m2[t2]++
							m3[t2] = append(m3[t2], m3[t1][j])
						}
						m3[t1] = make([]int32, 0)
						if m2[t2] > m2[max] {
							max = t2
						}
					}
				}
			} else {
				m1[queries[i][1]] = m1[queries[i][0]]
				m3[m1[queries[i][0]]] = append(m3[m1[queries[i][0]]], queries[i][1])
				m2[m1[queries[i][0]]]++
				if m1[queries[i][0]] != max && m2[max] < m2[m1[queries[i][0]]] {
					max = m1[queries[i][0]]
				}
			}
		} else if m1[queries[i][1]] > 0 {
			m1[queries[i][0]] = m1[queries[i][1]]
			m3[m1[queries[i][1]]] = append(m3[m1[queries[i][1]]], queries[i][0])
			m2[m1[queries[i][1]]]++
			if m1[queries[i][1]] != max && m2[max] < m2[m1[queries[i][1]]] {
				max = m1[queries[i][1]]
			}
		} else {
			m1[queries[i][0]] = idx
			m1[queries[i][1]] = idx
			m3[idx] = make([]int32, 0)
			m3[idx] = append(m3[idx], queries[i][0])
			m3[idx] = append(m3[idx], queries[i][1])
			m2[idx] += 2
			if m2[idx] > m2[max] {
				max = idx
			}
			idx++
		}
		fmt.Println(m1)
		fmt.Println(m3)
		fmt.Println("-----")
		ret[i] = int32(m2[max])
	}
	return ret
}
