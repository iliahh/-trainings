package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var result, value int

	for i := range n { //for i := 0; i < n; i++ -> for i := range n
		fmt.Scan(&value)

		if i == 0 {
			result = value // превое число просто записываем
		} else if i%2 == 1 {
			result -= value // нечётный индекс - вычитание
		} else {
			result += value // чётный индекс - прибавляем
		}
	}
	fmt.Println(result)
}
