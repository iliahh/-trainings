package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	dx := int(math.Abs(float64(x1 - x2))) // math.Abs нужен, чтобы получить модуль разницы координат,
	dy := int(math.Abs(float64(y1 - y2))) // т.к. направление не важно.
	// Ход коня возможен только если разность по координатам - это (2,1) или (1,2).
	// Остальные случаи - это не буква "Г", значит, нельзя.
	if (dx == 2 && dy == 1) || (dx == 1 && dy == 2) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	} // 3ms 2.00Mb
}
