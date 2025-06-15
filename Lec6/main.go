package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		for init; condition; post {
			init - блок инициализации переменных цикла
			condition - условие (если верно - тело цикла выполняется, если нет - цикл завершается)
			post - изменение переменной цикла (инкрементарное действие, декрементарное действие)
		}
	*/
	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}
	//Важный момент: в качестве init может быть использовано выражение присваивания ТОЛЬКО через :=

	//break - команда, прерывающая текущее выполнение тела цикла и передающая управление инструкциям, следующим за циклом
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	//continue - команда, прерывающая текущее выполнение тела цикла и передающая управление СЛЕДУЮЩЕЙ ИТЕРАЦИИ ЦИКЛА
	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with CONTINUE")

	//Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идее выше треугольник")

	//Иногда бывает плохо. С лейблами получше. Лейблы - это синтаксический сахар
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer //Хочу, чтобы вообще все циклы (внешние тоже) остановились
			}
		}
	}
	//Модификации цикла for.

	//1. Классический цикл while do
	var loopVar int = 0
	// while (loopVar <10){
	// 	....
	// 	LoopVar++
	// }
	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}

	//2. Классический бесконечный цикл
	var password string
outer2: //так лучше лейбл не называть, типо "2" не ставить
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
			continue
		} else {
			fmt.Println("Password Accepted")
			break outer2
			//return - типа ТРУ хакеры делают это
		}
	}

	//3. Цикл с множественными переменными цикла
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 { //x++, y++ -> x++; y++;
		fmt.Printf("%d+%d=%d\n", x, y, x+y)
	}
}
