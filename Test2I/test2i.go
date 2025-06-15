package main

import "fmt"

func main() {
	// var n int
	// var numbers []int

	// fmt.Scan(&n)
	// count := 0

	// for i := 1; i <= n; i++ {
	// 	if n%i == 0 {
	// 		count++
	// 		numbers = append(numbers, i)
	// 	}
	// }
	// for _, num := range numbers {
	// 	fmt.Printf("%d ", num)
	// }
	// if count <= 2 {
	// 	fmt.Printf("\nACHTUNG")
	// }
	var n, count int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			count++
		}
	}

	if count == 2 { // if count <= 2 - это для граничных случаев (1, 0 и т.д.)
		fmt.Printf("\nACHTUNG")
	}
}
