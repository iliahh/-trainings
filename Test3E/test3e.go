package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	runeSlice := []rune(word) //Записываем в слайс рун, чтобы выводилась кириллица

	for i := 0; i < len(runeSlice); i += 2 { //Делаем цикл с шагом в 2 и пределом в виде длины слайса (длина слова), куда пихаем необходимый вывод
		fmt.Printf("%c%c%c", runeSlice[i], runeSlice[i], runeSlice[i])
	}
	fmt.Println() //Добавлен перевод строки для лучшей читаемости
}
