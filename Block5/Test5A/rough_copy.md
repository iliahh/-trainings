1) Первая версия:
package main

import "fmt"

//Функция вычисления факториала
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

//Функция рассчёта числа сочетаний
func combination(n, m int) int {
	if m > n {
		return 0 // Нельзя выбрать больше, чем есть
	}
	return factorial(n) / (factorial(m) * factorial(n-m))
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	result := combination(n, m)
	fmt.Printf("%d\n", result)
}

2) Вторая версия:
package main

import (
	"fmt"
)

// Используем int64 для больших значений
func factorial(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func combination(n int64, m int64) int64 {
	if m > n {
		return 0
	}
	return factorial(n) / (factorial(m) * factorial(n-m))
}

func main() {
	var n, m int64
	fmt.Scan(&n)
	fmt.Scan(&m)

	fmt.Println(combination(n, m))
}

3) Третья версия:
package main

import (
	"fmt"
	"math/big"
)

// Функция вычисления факториала
func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

// Функция расчёта числа сочетаний
func combination(n, m int64) *big.Int {
	if m > n {
		return big.NewInt(0)
	}
	num := factorial(n)
	den1 := factorial(m)
	den2 := factorial(n - m)

	den := new(big.Int).Mul(den1, den2)
	return new(big.Int).Div(num, den)
}

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	result := combination(n, m)
	fmt.Println(result)
}