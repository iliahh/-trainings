1) Первый код:
package main

import (
	"fmt"
)

func main() {
	var word string

	fmt.Scan(&word)

	//Проверка на валидность вводимых символов, нужны "о" и "р"
	for _, ch := range word {
		if ch != 'о' && ch != 'р' && ch != 'О' && ch != 'Р' {
			return
		}
	}

	runeSlice := []rune(word)

	//Обработка случая, когда ничего не вводится
	if len(runeSlice) == 0 {
		return
	}

	maxCount := 1
	currentCount := 1

	for i := 1; i < len(runeSlice); i++ {
		if runeSlice[i] == 'о' && runeSlice[i] == runeSlice[i-1] {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 1
		}
	}
	fmt.Println(maxCount)
}
