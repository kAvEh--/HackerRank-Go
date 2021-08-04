package main

import (
	"fmt"
	"math/big"
)

func main() {
	extraLongFactorials(25)
}

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) {
	// Write your code here
	f := big.NewInt(1)
	for i := 2; i <= int(n); i++ {
		t := big.NewInt(int64(i))
		f.Mul(f, t)
	}
	fmt.Println(f)
}
