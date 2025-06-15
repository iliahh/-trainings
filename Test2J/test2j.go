package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := -n; i <= n; i++ { //Для всех i от -n до n включительно с шагом в единицу
		val := i * i //Ищем квадрат каждого числа в промежутке [-n;+n]
		fmt.Printf("Квадрат числа %d равен %d\n", i, val)
	}
}
