package main

import (
	"fmt"
)

func main() {
	// var a, b, c string
	// fmt.Scan(&a, &b, &c)
	// fmt.Println(c + b + a) //3ms 2.12Mb кривое решение, мне не нравится

	// var a, b, c int
	// fmt.Scan(&a, &b, &c)
	// fmt.Print(c, b, a) // Не работает по условию задачи

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d%d%d", c, b, a) //Здесь работаю с числами, а не костылем из строки
	//3ms 2.00Mb

	// var a, b, c int
	// fmt.Scan(&a, &b, &c)
	// result := strconv.Itoa(c) + strconv.Itoa(b) + strconv.Itoa(a)
	// //В этом решении преобразовываем в строки и склеиваем числа
	// fmt.Print(result) //3ms 2.13Mb
}
