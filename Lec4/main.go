package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//BooLean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	//Numerics. Integers
	//int8, int16, int32, int64, int - цифры тут в битах
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	//Вывод типа через %T форматирование:
	fmt.Printf("Type is %T\n", a)
	//Узнаем сколько байт занимает переменная типа int (получаем 8 байт, то есть 64 бита):
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент 1. При использовании короткого объявления тип для целого числа int платформо-зависимый.
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент 2. Используйте явное приведение типов при необходимости, если уверены, что не произойдёт коллизии.
	var first32 int32 = 12
	var second64 int64 = 13
	//fmt.Println(first32 + int32(second64))
	fmt.Println(int64(first32) + second64)

	//Эксперимент 3. Если проводятся арифметические операции
	// над Int и intX, то обяхательно нужно использовать механизм приведения. Т.к. int != int 64
	var third64 int64 = 1000000
	var fourthInt int = 12000
	fmt.Println(third64 + int64(fourthInt))

	//Явное приведение нужно при использовании разной битности int.

	// + - * / %

	//Аналогичным образом устроены uint8, uint16, uint32, uint64, uint
	//var varA uint =10

	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("Type of a %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

	//Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 332i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings. Строка - это набор БАЙТ (слайс), (concat - конкатизация)
	name := "Федя"
	lastname := "Pukich"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Lenght of string:", name, len(name)) //Функция len() возвращает количество элементов в наборе
	fmt.Println("Amount of cahrs:", utf8.RuneCountInString(name))
	//Rune - руна, это один utf-ный символ.
	//Поиск подстроки в строке:
	totalString, subString := "ABCDEDFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' //Для инициализации руны символьным значением используйте - ''
	var thirdRune rune = 234542
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)
	//Rune - это обвесок над int32 и под капотом там лежит нормальная цифорка для последующей работы
	//Для каждого кирилического символа нужно два байта, каждая строка это по умолчанию набор байт!

	//"A" < "abcd"
	//fmt.Println("abcd" > "a") - true, false при > "qwer"
	fmt.Println(strings.Compare("abcd", "a")) // -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte // alias uint8
	fmt.Println("Byte:", aByte)
}
