package main

import (
	"fmt"
)

// Функция вычисления факториала m
func factorial(m int) int {
	result := 1
	for i := 2; i <= m; i++ {
		result *= i
	}
	return result
}

// Функция расчёта числа сочетаний
func combination(n, m int) int {
	if m > n {
		return 0 //Нельзя выбрать больше, чем есть
	}

	//C(n, m) = n * (n - 1) * ... * (n - m + 1) / m!
	num := 1
	for i := range m {
		num *= (n - i)
	}

	den := factorial(m) //m!
	return num / den
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(combination(n, m))
}
