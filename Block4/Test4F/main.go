package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	all := make([]string, m+n)

	for i := 0; i < m+n; i++ {
		fmt.Scan(&all[i])
	}

	counts := make(map[string]int)
	//map[]
	for _, name := range all { //итерируем по слайсу all и записываем все строки из него мапу
		counts[name]++ //и подсчитываем их количество
	}
	//map[Иванов:2 Николаев:1 Петров:1 Сидоров:2]

	studs := 0
	for _, count := range counts { //итерируем по мапе, и если значение у элемента единица,
		if count == 1 { //а не двойка, тогда подсчитываем в переменную studs
			studs++
		}
	}

	fmt.Println(studs)
}
