package main

import (
	"fmt"
)

// НОД (нужен для сокращения дроби)
func NOD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// НОК (нужен для приведения к общему знаменателю)
func NOK(a, b int) int {
	return a * b / NOD(a, b)
}

func main() {
	var n int
	fmt.Scan(&n)

	var totalNum, totalDen int
	for i := range n {
		var num, den int
		fmt.Scan(&num)
		fmt.Scan(&den)

		if i == 0 {
			// первую дробь просто записываем как базовую
			totalNum = num
			totalDen = den
		} else {
			// приводим к общему знаменателю
			lcmDen := NOK(totalDen, den)
			totalNum = totalNum*(lcmDen/totalDen) + num*(lcmDen/den)
			totalDen = lcmDen
		}
	}

	// Сокращаем итоговую дробь
	g := NOD(totalNum, totalDen)
	fmt.Printf("%d/%d\n", totalNum/g, totalDen/g)
}
