package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello world!"
	fmt.Println(name) //Hello world!

	//1. Строка - это байтовый слайс со своими особенностями при обращении к нижележащему массиву
	word := "Тестовая строка"       //"Sample word"
	fmt.Printf("String %s\n", word) //String Sample word

	// Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-ти ричного байта
	} //Bytes: 53 61 6d 70 6c 65 20 77 6f 72 64
	fmt.Println()

	//Каким образом получать доступ к отдельно стоящим символам?
	fmt.Printf("Characters: ") //Characters: S a m p l e   w o r d //Characters: Ð ¢ Ð µ Ñ  Ñ  Ð ¾ Ð ² Ð ° Ñ    Ñ  Ñ  Ñ  Ð ¾ Ð º Ð °
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
	fmt.Println()

	//2. Строки в GO - хранятся как наборы UTF-8 символов. Каждый символ, вообще говоря,
	// может занимать больше чем один байт.

	//3. Руна (Rune) - это стандартный встроенный тип в GO (alias над int32), позволяющий хранить
	//единый неделимый UTF символ ВНЕ ЗАВИСИМОСТИ ОТ ТОГО, сколько байт он занимает

	fmt.Printf("Runes: ")     // Runes: S a m p l e   w o r d // Runes: Т е с т о в а я   с т р о к а
	runeSlice := []rune(word) // Преобразование слайса байт к слайсу рун []byte{sliceRune}
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()

	//4. Итерирование по строке с использованием рун
	for id, runeVal := range word { // id - это индекс байта, С КОТОРОГО НАЧИНАЕТСЯ РУНА runeVal
		fmt.Printf("%c starts at position %d\n", runeVal, id)
	}
	// Т starts at position 0
	// е starts at position 2
	// с starts at position 4
	// т starts at position 6
	// о starts at position 8
	// в starts at position 10
	// а starts at position 12
	// я starts at position 14
	//   starts at position 16
	// с starts at position 17
	// т starts at position 19
	// р starts at position 21
	// о starts at position 23
	// к starts at position 25
	// а starts at position 27

	//5. Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43} // Исходное представление байтов
	myStr := string(myByteSlice)
	fmt.Println(myStr) //@ABC

	//Представление в десятичном формате
	myDecimalBytesSlice := []byte{100, 101, 102, 103} // Синтаксический сахар - можно использовать десятичное представление байтов
	myDecimalStr := string(myDecimalBytesSlice)
	fmt.Println(myDecimalStr) //defg

	//6. Создание строки из слайса рун
	// Руны как hex
	runHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRune := string(runHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRune) //Form Runes(hex): EFGH
	// Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 'y', 'a'} // '' - таким образом обозначается руна
	myStrFormRuneLiterals := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFormRuneLiterals) //From Runes(literals): Vasya

	//Чтобы убедиться, что это действительно строка:
	fmt.Printf("%s and %T\n", myStrFormRuneLiterals, myStrFormRuneLiterals) //Vasya and string

	//7. Длина и ёмкость строки
	//Длина Len() - количество байт в слайсе
	fmt.Println("Length of Вася:", len("Вася"), "bytes") //Length of Вася: 8 bytes
	//Длина RuneCounter - количество !рун!
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes") //Length of Вася: 4 runes
	//По факту работает как цикл из 4 пункта
	//Вычисление ёмкости строки - бессмысленно, т.к. строка базовый тип
	fmt.Println(cap([]rune("Вася"))) //4

	//8. Сравнение строк == и !=. Начиная с GO v1.6
	word1, word2 := "Вася", "Петя"
	if word1 != word2 { //Not equal
		// if word1 != word2 //Equal
		// if word1 > word2 //Not equal определение идёт вообще по байтам
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	//9. Конкатенация (при конкатенации создаётся новая строка на основе старых двух)
	word3 := word1 + word2
	fmt.Println(word3) //ВасяПетя

	//10. Строитель строк (java -> StringBuilder)
	firtsName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firtsName, secondName)
	fmt.Println("Fullname:", fullName) //Fullname: Alex #### Johnson

	//11. Строки не изменяемые
	//fullName[0] = "Q" //Что бы тут не использовалось, всегда выдаст ошибку
	//Lec10\main.go:116:2: cannot assign to fullName[0] (neither addressable nor a map index expression)

	//12. А слайсы изменяемые :)
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	fullName = string(fullNameSlice)
	fmt.Println("String mutation:", fullName) //String mutation: Flex #### Johnson

	//13. Сравнение рун
	if 'Q' == 'Q' { // if 'Q' == 'q' //Runes not equal
		fmt.Println("Runes equal")
	} else {
		fmt.Println("Runes not equal")
	}
	//Runes equal

	//14. Где живут полезные методы работы со строками?
	//import "strings"
}
