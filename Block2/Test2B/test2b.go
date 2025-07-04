package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		// Линейное уравнение bx + c = 0
		if b == 0 {
			fmt.Print("корней нет") // даже если c == 0, по условию решений нет
		} else {
			fmt.Print("один корень") // b ≠ 0 → есть решение
		}
	} else {
		// Квадратное уравнение ax² + bx + c = 0
		D := b*b - 4*a*c
		if D > 0 {
			fmt.Print("два корня")
		} else if D == 0 {
			fmt.Print("один корень")
		} else {
			fmt.Print("корней нет")
		}
	}
	// a = 2, b = 5, c = 2 -> дискриминант > 0 -> два корня
	// a = 2, b = 4, c = 2 -> дискриминант = 0 -> один корень
	// a = 0, b = 6, c = 0 -> не квадратное -> корней нет
	// a = 2, b = 1, c = 5 -> дискриминант < 0 -> корней нет
}
