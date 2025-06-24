package main

import (
	"fmt"
	"math"
)

// canKnightMove проверяет, может ли конь сделать ход из (x1, y1) в (x2, y2)
func canKnightMove(x1, y1, x2, y2 int) bool {
	// Считаем абсолютную разницу по координатам
	dx := int(math.Abs(float64(x1 - x2)))
	dy := int(math.Abs(float64(y1 - y2)))

	// Проверяем одно из двух правил хода конем: 2 по одной оси, 1 по другой
	return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
}

func main() {
	var x1, y1, x2, y2 int

	// Считываем 4 числа — координаты начальной и целевой клетки
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	// Используем функцию, чтобы решить задачу
	if canKnightMove(x1, y1, x2, y2) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
