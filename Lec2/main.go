package main

import (
	"fmt"
	"math"
)

func main() {
	//Простейший вывыод на консоль
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//Форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
	///////////////////////////
	///////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 4
	fmt.Println("Age after initialization:", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 191
	fmt.Println("My height is:", height)

	//В чём "полустрогость" типизации? Можно опускать тип переменной, компилятор подставит сам
	var uid = 12345
	fmt.Println("My uid:", uid)
	//fmt.Printf("%T\n", uid)
	//Декларирование и инициализация переменных одного типа (множественный случай).
	var firstVar, secondVar = 20, 30 //тут было int = 20, 30, но тут тоже можно опустить тип переменной.
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)
	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge         = 42 //тут было int = 42, но тут тоже можно опустить тип переменной.
		personUID  int
	)
	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)

	//Немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b) //если мы не задаём тип переменной, то Go делает это сам на своё усмотрение.
	//Лучше ставить тип переменной и не лениться, ставим блок и в нем прописываем типы переменных.

	//Немного хорошего: Повторное декларирование переменной приводит к ошибке компиляции
	// если мы задекларировали через блок var переменную,
	// то заново её задекларировать мы не можем.
	//var a = 200 - ломает компиляцию, a = 200 - просто перезаписывает переменную.

	//Коротная декларация (короткое объявление) Требует моментальной инициализации какого-то значения.
	count := 10
	fmt.Println("Count", count)
	count = count + 1
	fmt.Println("Count", count)

	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg) //Если в блоке присваивания через := есть хотя бы одна новая переменная, то старым будет присвоено новое значение, это особенность языка.

	//Пример
	width, lenght := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.2f\n", math.Min(width, lenght))
}
