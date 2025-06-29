package main

import "fmt"

//1. Явные функции (в принципе любая функция в Go) - являются
//экземпляром 1-го уровня (функцию можно присваивать в переменную, её можно передавать в
//качестве параметра и возвращать из других функций)

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func mult(a, b int) int {
// 	return a * b
// }

// func calcAndReturnValidFunc(command string, a int, b int) int {
// 	if command == "addition" {
// 		return func(a, b int) int { return a + b }(a, b)
// 	} else if command == "substraction" {
// 		return func(a, b int) int { return a - b }(a, b)
// 	} else {
// 		return func(a, b int) int { return a * b }(a, b)
// 	}
// }

//2. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

// func sample() func() {
// body
// }

//3. Функция как параметр в другой функции
func recieveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

//Пример:
func add(a, b int) int {
	return a + b
}

func main() {

	var command string
	// command = "addition"
	command = "substraction"
	// res := calcAndReturnValidFunc(command, 10, 20)
	res := calcAndReturnValidFunc(command)
	// fmt.Println("Result with command: ", command, "value:", res) //Result with command:  addition value: 30
	fmt.Println("Result with command: ", command, "value:", res(10, 20)) //Result with command:  addition value: 30

	fmt.Println(res(30, 40)) //Result with command:  substraction value: -10 //-10

	//4. Типи функции
	fmt.Printf("Type of func is %T\n", res)                                      //Type of func is func(int, int) int
	fmt.Printf("Type of calcAndReturnValidFunc is %T\n", calcAndReturnValidFunc) //Type of calcAndReturnValidFunc is func(string) func(int, int) int

	//5. Тип функции в Go определяется как входными параметрами, так и выходными
	fmt.Println("recieveFuncAndReturnValue(add):", recieveFuncAndReturnValue(add)) //recieveFuncAndReturnValue(add): 300
	fmt.Println(recieveFuncAndReturnValue(func(a, b int) int {                     //90000
		return a*a + b*b + 2*a*b
	}))
}
