package main

import "fmt"

//1. Явная функция - локально-определённый блок кода, имеющий имя (ЯВНОЕ ОПРЕДЕЛЕНИЕ)
//Функцию необходимо определить + функцию необходимо вызвать
//2. Сигнатура функций и их определение
// func functionName(params type) typeReturnValue {
// 	//body
// }

//3. Простейшая функция (определить функцию можно как до момента её вызова в функции main,
//так и в любом месте пакета, главное чтобы она была определена в принципе где-нибудь)
func add(a int, b int) int {
	result := a + b
	return result
}

func main() {
	fmt.Println("Hellow world") //Hellow world
	//4. Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)                   //Result of add(10, 20): 30
	fmt.Println("Result of mult(1, 2, 3, 4):", mult(1, 2, 3, 4)) //Result of mult(1, 2, 3, 4): 24

	per, area := rectangleParameters(10.5, 10.5)
	//Если один параметр не нужен:
	newPer, _ := rectangleParameters(10, 10)
	fmt.Println("Area of rect(10.5, 10.5):", area)     //Area of rect(10.5, 10.5): 110.25
	fmt.Println("Perimeter of rect(10.5, 10.5):", per) //Perimeter of rect(10.5, 10.5): 42
	fmt.Println("NewPer:", newPer)                     //NewPer: 40
	secondPer, secondArea := namedReturn(10, 20)
	fmt.Println(secondPer, secondArea) //60 200

	emptyReturn(10) //resEmpty := emptyReturn(10) - ошибка

	helloVariadic(10, 20, 30, 40, 50, 60, 60) //[10 20 30 40 50 60 60]
	helloVariadic(10, 20)                     //[10 20]
	helloVariadic(10)                         //[10]
	helloVariadic()                           //[]

	someStrings(2, 10, "Bob", "Alex", "Vito")
	//Parameter: 2
	//Parameter: 10
	//Result concat: BobAlexVito

	sum1 := sumVariadic(10, 30, 40, 50, 60)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...) //Или сразу sum2 := sumVariadic([]int{10, 20, 30}...)
	fmt.Println(sum1, sum2)             //190 60

	fmt.Println(sumSlice([]int{30, 40, 50, 60, 80, 90, 100})) //450 - заморочка и гадость, лучше так не делать
	fmt.Println(sumSlice(sliceNumber))                        //60

	//12. Анонимная функция (синтаксис)
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println("Anon:", anon(20, 30))

	// go func(){

	// }

	fmt.Println("BigFunction(10, 20):", bigFunction(10, 20)) //BigFunction(10, 20): 31
}

//4. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

//5. Возврат больше чем одного значения (returnType1, returnType2...)
func rectangleParameters(lenght, width float64) (float64, float64) {
	var perimeter = 2 * (lenght + width)
	var area = lenght * width

	return perimeter, area
}

//6. Именованный возврат значений
func namedReturn(a, b int) (perimeter, area int) {
	//:= не используется, т.к. обозначили perimeter выше в функции, у этой переменной уже есть ноль,
	// и мы его перезаписываем ниже обычным =
	perimeter = 2 * (a + b)
	area = a * b
	return //Не нужно указывать возвращаемые переменные
}

// func sample(a int) (int, int) {
// 	var res1 = a * a
// 	var res2 = a + a
// 	return res1, res2
// }

//7. При вызове оператора return функция прекращает своё выполнение и возвращает что-то
func funcWithReturn(a, b int) (int, bool) {
	if a > b { //если 5 и 7, то до второго if код не дойдет
		return a - b, true
	}

	if a == b { //если 10 и 7, то код уйдет до return b
		return a, true
	}

	return 0, false
}

//8. Что вернётся в случае, если return в принципе не указан (или он пустой)
func emptyReturn(a int) {
	fmt.Println("I'm emptyReturn with parameter:", a) //I'm emptyReturn with parameter: 10
	//return - можно убрать
}

//9. Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Println(a)
	fmt.Printf("value %v and type %T\n", a, a) //value [] and type []int - это слайс
}

//10. Смешение параметров с variadic (
//	1. Континуальный параметр всегда самый последний
//	2. Variadic параметр - на всю функцию один (для удобочитаемости)
//	)
func someStrings(a, b int, words ...string) { //func someStrings(words ...string, a int, b int)
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

//11. Передача слайса или использование variadic parameters?
func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func sumSlice(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

//13. Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg) //Извращенство, короче
}
