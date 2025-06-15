package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	//Нужно использовать float64 для math.Pow, поэтому преобразовываем int в float64 и обратно
	sum := float64(a + b)
	result := int(math.Pow(sum, 2))
	fmt.Println(result) //3ms 2.00Mb
}
