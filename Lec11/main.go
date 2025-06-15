package main

import "fmt"

func main() {
	//1. Map - это набор пар ключ:значение. Инициализация пустой мапы
	mapper := make(map[string]int)    //Из строки в целое число
	fmt.Println("Empty map:", mapper) //Empty map: map[]

	//2. Добавление пар в существующую мапу
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs:", mapper) //Mapper after adding pairs: map[Alice:24 Bob:25]

	//3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}
	newMapper["Jo"] = 3000
	fmt.Println("New Mapper:", newMapper) //New Mapper: map[Alice:1000 Bob:1000]
	//New Mapper: map[Alice:1000 Bob:1000 Jo:3000]

	//4. Что может быть ключом в мапе?
	//4.1 ВАЖНО: Мапа НЕ УПОРЯДОЧЕНА В GO
	//4.2 КЛЮЧОМ В МАПЕ МОЖЕТ БЫТЬ ТОЛЬКО СРАВНИМЫЙ ТИП (==, !=)

	//5. Нулевое значение для map
	// var mapZeroValue map[string]int // mapZeroValue == nil
	// mapZeroValue["Alice"] = 12

	//6. Получение элементов из map
	//6.1 Получение элемента, который представлен в мапе
	testPerson := "Alice"
	fmt.Println("Salary of testPerson:", newMapper[testPerson]) //Salary of testPerson: 1000
	//6.2 Получение элемента, который НЕ представлен в мапе
	testPerson = "Derek"
	fmt.Println("Salary of new testPerson:", newMapper[testPerson]) //При обращении по несуществующему ключу, новая пара не добавляется
	//Salary of new testPerson: 0
	fmt.Println(newMapper)

	//7. Проверка вхождения ключа
	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	fmt.Println("Den and value:", employee["Den"]) //Den and value: 0
	fmt.Println("Jo and value:", employee["Jo"])   //Jo and value: 0

	//7.1 При обращении по ключу - возвращается 2 значения
	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo and value:", value) //Den and value: 0 //Jo does not exists in map
	} else {
		fmt.Println("Jo does not exists in map")
	}

	//8. Перебор элементов мапы
	fmt.Println("==============================")
	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}
	// Den and value 0
	// Alice and value 0
	// Bob and value 0

	//9. Как удалять пары
	//9.1 Удаление существующей пары
	fmt.Println("Before deleting:", employee) //Before deleting: map[Alice:0 Bob:0 Den:0]
	delete(employee, "Den")
	fmt.Println("After first deleting:", employee) //After first deleting: map[Alice:0 Bob:0]

	//9.2 Удаление несуществующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") //Эта функция delete ничего не возвращает и это ОЧЕНЬ ДОРОГАЯ ОПЕРАЦИЯ
	}
	fmt.Println("After second deleting:", employee)

	//10. Количество пар == длина мапы
	fmt.Println("Pair amount in map:", len(employee)) //Pair amount in map: 2

	//11. Мапа (как и слайс) ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words map:", words)       //words map: map[Three:Три Two:Два]
	fmt.Println("newWords map:", newWords) //newWords map: map[Three:Три Two:Два]

	//12. Сравнение мап
	//12.1 Сравнение массивов (массив можно использовать в качестве ключа в мапе)
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	//if [3]int{1, 2, 3} == [3]int{3, 4, 5} //Not equal
	//if [3]int{1, 2, 3} == [3]int{1, 2, 3} //Equal

	//12.2 Сравнение слайсов
	// if []int{1, 2, 3} == []int{1, 2, 3} {
	// 	fmt.Println()
	// }
	// Слайсы нельзя сравнивать друг с другом из-за того,
	// что тип ссылочный, и можно сравнить только равенство с nil

	//12.3 Сравнение мап
	// aMap:=map[string]int{
	// 	"a":1,
	// }
	// bMap:=map[string]int{
	// 	"b":1,
	// }
	// if aMap == bMap{
	// ...
	// }
	//map can only be compared to nil
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero value of map bMap")
	}
	//Zero value of map bMap

	//13. Грустное следствие
	//Если мапа/слайс являются составляющими какой-либо структуры, то структура автоматически не сравнима

}
