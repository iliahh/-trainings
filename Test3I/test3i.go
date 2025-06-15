package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin) //Количество задач, из которых будут создавать напоминания

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	words := make([]string, n) //Создаём слайс для списка дел

	for i := range n { //Обозначаем цикл считывания n элементов
		scanner.Scan()
		words[i] = scanner.Text() //Считываем n элементов
	}

	scanner.Scan() //Считываем количество задач, из которых будет собрано напоминание
	m, _ = strconv.Atoi(scanner.Text())

	selected := []string{} //Создаём пустышку для напоминалки дел
	for range m {          //Обозначаем цикл, в котором обозначатся m элементов
		scanner.Scan()
		index, err := strconv.Atoi(scanner.Text()) //Считываем индексы(номера) дел, которые будут выводиться
		if err != nil || index < 1 || index > n {  //Количество дел в напоминалке не превышает максимальное заданное и не может быть меньше одного
			fmt.Println(index)
			continue
		}
		selected = append(selected, words[index-1]) //Теперь в пустышку запишем наши дела из первого слайса
	}

	for _, task := range selected { //Игнорирование id в "range based for" цикле
		fmt.Println(task) //Вывод самих дел в напоминалке
	}
}
