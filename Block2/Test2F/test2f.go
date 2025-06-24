package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print(0)
		return
	}

	var bits []int
	for n > 0 {
		bits = append(bits, n%2)
		n = n / 2
	}

	for i := len(bits) - 1; i >= 0; i-- {
		fmt.Print(bits[i])
	}
}
