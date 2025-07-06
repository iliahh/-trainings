package main

import (
	"fmt"
	"sort"
)

type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

// Отсортировать слова и распечатать их
func (sl *SmartList) GetAnswer() {
	sort.Slice(sl.words, func(i, j int) bool { //sort.Slice(что_сортируем, как_сравниваем)
		if len(sl.words[i]) == len(sl.words[j]) { //Условие сравнения одинаковых по длине слов
			return sl.words[i] < sl.words[j] //Возвращение отсортированных по алфавиту слов одинаковой длины
		}
		return len(sl.words[i]) < len(sl.words[j]) //Возвращение отсортированных по длине слов
	}) //Возвращаются в слайс, который уже является изменённым

	for _, word := range sl.words { //Далее итерируем по изменённому слайсу (в слайсе изменён порядок элементов, меняется он не в копии, т.к. sl.words указывает на оригинальный слайс, что был передан в New())
		fmt.Println(word) //Выводим отсортированные результаты
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	wordSlice := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&wordSlice[i])
	}

	r := New(wordSlice) //Закидываю слайс n слов в конструктор
	r.GetAnswer()       //Вызываю метод для получения отсортированного вывода слов
}
