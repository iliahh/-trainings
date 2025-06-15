package main

import "fmt"

func main() {
	var a, b, n, val, sum int
	fmt.Scan(&n)

	var slice []int

	for i := 1; i <= n; i++ {
		fmt.Scan(&val)
		slice = append(slice, val)
	}

	fmt.Scan(&a, &b)

	for i := a - 1; i < b; i++ {
		sum += slice[i]
	}
	fmt.Println(sum)
}
