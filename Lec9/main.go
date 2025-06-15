package main

import "fmt"

func main() {
	//1. Слайсы (они же - срезы)
	// Слайс - это динамическая обвязка над массивом.
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками!
	//Добавь от нулевой до второй не включительно элементы из startArr
	fmt.Println("Slice[0:2]:", startSlice) //Slice[0:2]: [10 20]
	// Создали слайс, основываясь уже на существующем массиве

	//2. Создание слайса без явной инициализации массива
	secondSLice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSLice) //SecondSlice: [15 20 30 40]

	//3. Изменение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4] //Это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++ //В итоге добавляем по единичке к каждому элементу среза,
		// сделанного из 1-3 элементов originArr := [...]int{30, 40, 50, 60, 70, 80}
	}
	fmt.Println("OriginArr:", originArr)   //OriginArr: [30 41 51 61 70 80]
	fmt.Println("FirstSlice:", firstSlice) //FirstSlice: [41 51 61]

	//4. Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	// Before modifications: Arr: [30 41 51 61 70 80] fSlice: [30 41 51 61 70 80] sSlice: [51 61 70]
	// After modifications: Arr: [30 41 51 63 70 80] fSlice: [30 41 51 63 70 80] sSlice: [51 63 70]

	//5. Срез как встроенный тип
	// type slice struct{
	// 	Length int
	// 	Capacity int
	// 	ZeroElement *byte
	// }

	//6. Длина и ёмкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("Slice:", wordsSlice, "Length:", len(wordsSlice), "Capacity:", cap(wordsSlice))
	//Slice: [one two three] Length: 3 Capacity: 3
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("Slice:", wordsSlice, "Length:", len(wordsSlice), "Capacity:", cap(wordsSlice))
	//Slice: [one two three four] Length: 4 Capacity: 6
	//Capacity (cap) или ёмкость слайса - это значение, показывающее СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ
	//можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ.

	//Допустим у нас есть срез на 3 элемента (инициализировали без явного указания массива)
	//Компилятор при создании этого среза СНАЧАЛА создал массив ровно на 3 элемента
	//После этого компилятор вернул адрес, где этот массив живёт, срез запомнил этот адрес и теперь
	//ссылается на него. Потом начинаем деформировать слайс (увеличим длину / увеличим количество элементов)
	//Проблема - в нижележащем массиве (на котором основан слайс) всего 3 ячейки. Что делать?
	//Компилятор ищет в памяти место для массива размера 3*2 (в общем случае n*2, где n - первоначальный размер)
	//После того, как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	//старые 3 элемента на свои позиции. На 4-ую позицию мы добавляем новый элемент
	//После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив под 6 элементов.

	//Ёмкость всегда будет изменяться как n*2!
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%10 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
		// Current len: 2 Current cap: 2
		// Current len: 12 Current cap: 16
		// Current len: 22 Current cap: 32
		// Current len: 32 Current cap: 32
		// Current len: 42 Current cap: 64
		// Current len: 52 Current cap: 64
		// Current len: 62 Current cap: 64
		// Current len: 72 Current cap: 128
		// Current len: 82 Current cap: 128
		// Current len: 92 Current cap: 128
		// Current len: 102 Current cap: 128
		// Current len: 112 Current cap: 128
		// Current len: 122 Current cap: 128
		// Current len: 132 Current cap: 256
		// Current len: 142 Current cap: 256
		// Current len: 152 Current cap: 256
		// Current len: 162 Current cap: 256
		// Current len: 172 Current cap: 256
		// Current len: 182 Current cap: 256
		// Current len: 192 Current cap: 256
	}

	//Важно: после выделения памяти под новый массив, ссылки со старым будут перенесены в новый
	//Пример:
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) //В этот момент numSLice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)   // [1 2]
	fmt.Println(numSlice) // [1 2 3]
	//Упомянули алоцирование какое-то

	//7. Как создавать слайсы наиболее эффективно?
	// make() - это функция, позволяющая более детально создавать срезы
	sl := make([]int, 10, 15) //Длина не может быть > capacity
	// []int - тип коллекции
	// 10 - длина
	// 15 - ёмкость (capacity)
	//Сначала инициализируется arr = [15]int
	//Затем по нему делается срез arr[0:10]
	//После чего он возвращается
	fmt.Println(sl) //[0 0 0 0 0 0 0 0 0 0]

	//8. Добавление элементов в СРЕЗ
	// В массив нельзя добавить элементы, в срез можно!
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords) //myWords: [one two three]
	anotherSlice := []string{"four", "five", "six"}
	// for _, val := range anotherSlice {
	// 	myWords = append(myWords, val)
	// }
	//myWords = append(myWords, "four") Но есть более простой синтаксис
	myWords = append(myWords, anotherSlice...) //Фактически выглядит как myWords = append(myWords, "four", "five", "six")
	//fmt.Println("myWords:", myWords) //myWords: [one two three four]
	fmt.Println("myWords:", myWords) // myWords: [one two three four five six]
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords:", myWords) // myWords: [one two three four five six seven eight]

	//9. Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice) // [[1 2] [3 4 5 6] [10 20 30] []]
}
