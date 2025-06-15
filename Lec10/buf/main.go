package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() { //Команда захвата потока ввода и сохранения в буфер (захват идёт до символа окончания строки)
		name = input.Text() // Команда возвращения элементов, помещённых в буфер (отдаст string)
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}
	//input.Scan() - помещение в буфер
	//input.Text() - считывания из буфера
	//input.Bytes() - можно прям в байтах вернуть и далее поместить в слайс

	//Преобразование строкового литерала к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)        // Atoi - Anything to Int (именно int)
	fmt.Printf("%d is %T\n", numInt, numInt) //10 is int

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	//fmt.Println(numInt64 + numInt)
	//Lec10\buf\main.go:38:14: invalid operation: numInt64 + numInt (mismatched types int64 and int)
	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	// Но это 64-разрядное число будет без проблем ГАРАНТИРОВАНО ПЕРЕВОДИТЬСЯ К 32
	fmt.Println(numInt64, numFloat32)                            //10 10
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32)) //fmt.Printf("%.3f and %T\n", numFloat32, numFloat32) 10.000 and float64
	//10.000 and float32
}
