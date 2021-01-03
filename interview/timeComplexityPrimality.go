package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(primality(123))
}

// Complete the primality function below.
func primality(n int32) string {
	if big.NewInt(int64(n)).ProbablyPrime(0) {
		return "Prime"
	} else {
		return "Not prime"
	}

}
