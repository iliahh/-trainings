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

	maxCount := 0
	currentCount := 0

	//Ранее считались ТОЛЬКО подряд идущие "о", игнорировались одиночные,
	//но теперь счёт идёт всегда и просто обнуляется, когда прерывается серия
	for _, ch := range runeSlice {
		if ch == 'о' || ch == 'О' {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	fmt.Println(maxCount)
}
