package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	dens := make([]int, n)

	for i := range n {
		fmt.Scan(&nums[i])
		fmt.Scan(&dens[i])
	}

	// Считаем общий знаменатель в лоб: перемножаем всё
	totalDenom := 1
	for i := range n {
		totalDenom *= dens[i]
	}

	// Считаем новый числитель
	totalNum := 0
	for i := range n {
		mult := totalDenom / dens[i]
		totalNum += nums[i] * mult
	}

	// Упрощаем дробь
	g := NOD(totalNum, totalDenom)
	fmt.Printf("%d/%d\n", totalNum/g, totalDenom/g)
}

// простая функция НОД
func NOD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
