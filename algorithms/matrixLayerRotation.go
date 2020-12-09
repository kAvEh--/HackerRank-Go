package main

import "fmt"

func main() {
	data := [][]int32{{9718805, 60013003, 5103628, 85388216, 21884498, 38021292, 73470430, 31785927},
		{69999937, 71783860, 10329789, 96382322, 71055337, 30247265, 96087879, 93754371},
		{79943507, 75398396, 38446081, 34699742, 1408833, 51189, 17741775, 53195748},
		{79354991, 26629304, 86523163, 67042516, 54688734, 54630910, 6967117, 90198864},
		{84146680, 27762534, 6331115, 5932542, 29446517, 15654690, 92837327, 91644840},
		{58623600, 69622764, 2218936, 58592832, 49558405, 17112485, 38615864, 32720798},
		{49469904, 5270000, 32589026, 56425665, 23544383, 90502426, 63729346, 35319547},
		{20888810, 97945481, 85669747, 88915819, 96642353, 42430633, 47265349, 89653362},
		{55349226, 10844931, 25289229, 90786953, 22590518, 54702481, 71197978, 50410021},
		{9392211, 31297360, 27353496, 56239301, 7071172, 61983443, 86544343, 43779176}}
	matrixRotation(data, 40)
}

// Complete the matrixRotation function below.
func matrixRotation(matrix [][]int32, r int32) {
	rotate(matrix, len(matrix), len(matrix[0]), int(r))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func rotate(matrix [][]int32, m int, n int, r int) {
	c1 := (len(matrix) - m) / 2
	c2 := (len(matrix[0]) - n) / 2
	for j := 0; j < r%(2*(m+n-2)); j++ {
		rotateIt(matrix, c1, m, c2, n)
	}

	if m > 2 && n > 2 {
		rotate(matrix, m-2, n-2, r)
	}
}

func rotateIt(matrix [][]int32, d1 int, m int, d2 int, n int) {
	c2 := matrix[d1][d2+n-1]
	c3 := matrix[d1+m-1][d2]
	for i := m - 2; i >= 0; i-- {
		matrix[d1+i+1][d2] = matrix[d1+i][d2]
	}
	for i := 1; i < m; i++ {
		matrix[d1+i-1][d2+n-1] = matrix[d1+i][d2+n-1]
	}
	for i := 0; i < n-2; i++ {
		matrix[d1][d2+i] = matrix[d1][d2+i+1]
	}
	for i := n - 1; i > 1; i-- {
		matrix[d1+m-1][d2+i] = matrix[d1+m-1][d2+i-1]
	}
	matrix[d1][d2+n-2] = c2
	matrix[d1+m-1][d2+1] = c3
}
