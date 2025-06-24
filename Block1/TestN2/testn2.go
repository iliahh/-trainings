package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	dx := x1 - x2
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y2
	if dy < 0 {
		dy = -dy
	}

	if (dx == 2 && dy == 1) || (dx == 1 && dy == 2) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
