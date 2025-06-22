package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scan(&m) //m - число блюд, которое может приготовить столовая

	menu := make([]string, m) //Пустой слайс блюд для столовой с вводимым фиксированным числом

	for i := 0; i < m; i++ {
		fmt.Scan(&menu[i])
	}

	fmt.Scan(&n) //n - число дней, для которых есть списки блюд

	days := make([][]string, n) //n-блоков строчек для списков блюд на каждый день

	//Заполняем n-блок блюдами
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		dayMenu := make([]string, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&dayMenu[j])
		}
		days[i] = dayMenu //Добавляем в общий список
	}

	//Поиск всех блюд, которые хоть раз встречались
	difference := make(map[string]bool)
	for _, day := range days {
		for _, dish := range day {
			difference[dish] = true
		}
	}

	//Поиск блюд из menu, которые не встречались ни в одном дне
	var result []string
	for _, dish := range menu {
		if !difference[dish] {
			result = append(result, dish)
		}
	}

	sort.Strings(result) //Делаем порядок вывода алфавитным

	for _, dish := range result {
		fmt.Println(dish)
	}
}
