package main

import (
	"fmt"
)

func main() {
	fmt.Println(flippingBits(0))
}

// Complete the flippingBits function below.
func flippingBits(n int64) int64 {
	max := int64(4294967295)

	return max - n
}
