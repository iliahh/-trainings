package main

import (
	"fmt"
	"os"
)

func main() {
	// Открываем файл input.txt на чтение
	input, err := os.Open("input.txt") //"Block4/Test4C/input.txt"
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var n, k int
	var s string

	// Считываем данные из файла
	_, err = fmt.Fscan(input, &n, &k, &s)
	if err != nil {
		panic(err)
	}

	// Создаём файл output.txt для записи
	output, err := os.Create("output.txt") //"Block4/Test4C/output.txt"
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// Строим рамку
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if i == 0 || i == n-1 || j == 0 || j == k-1 {
				fmt.Fprint(output, s)
			} else {
				fmt.Fprint(output, " ")
			}
		}
		fmt.Fprintln(output)
	}
}
