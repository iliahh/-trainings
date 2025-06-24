package main

import "fmt"

func main() {
	// var a, b int
	// fmt.Scan(&a, &b)
	// p := 2*(a * b) + (b * 2)
	// s := a * b
	// fmt.Printf("Периметр: %d\nПлощадь: %d\n", p, s) - 3ms 2.13Mb
	var a, b int

	// Считываем две стороны прямоугольника
	fmt.Scan(&a)
	fmt.Scan(&b)

	// Вычисляем периметр и площадь
	perimeter := 2 * (a + b)
	area := a * b

	// Выводим результат
	fmt.Println("Периметр:", perimeter)
	fmt.Println("Площадь:", area) //3ms 2.12Mb
}
