package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil { // nil - аналог null, проверяем "есть ли ошибка?"
			break
		}
		if input == "" {
			break
		}
		fmt.Println(input)
	}
}
