1) Первая версия:
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		count := 2*i + 1
		for j := 0; j < count; j++ {
			fmt.Print("@")
		}
		fmt.Println()
	}
}