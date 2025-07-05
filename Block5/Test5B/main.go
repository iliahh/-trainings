package main

import "fmt"

type Sequence struct { //Добавляем структуру последовательности
	message string //В структуру добавляем поле типа string
}

func New(newMessage string) *Sequence { //Добавляем конструктор с входным новым сообщением и выходной последовательностью
	return &Sequence{newMessage} //Возвращаем адрес последовательности с новым сообщением
}

//Рассчитать максимальный стрик и вернуть это значение
func (s *Sequence) FindMax() int { //Добавляем метод для поиска стрика из "о"
	runeSlice := []rune(s.message)

	maxCount := 0
	currentCount := 0

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

	return maxCount
}

func (s *Sequence) Check() bool { //Метод на проверку валидности вводимых символов, нужны "о" и "р"
	for _, ch := range s.message {
		if ch != 'о' && ch != 'р' && ch != 'О' && ch != 'Р' {
			return false
		}
	}
	return true
}

func main() {
	var word string
	fmt.Scan(&word)

	s := New(word) //Вызов метода проверки валидности вводимого
	if !s.Check() {
		return
	}
	fmt.Println(s.FindMax())
}
