package main

import (
	"fmt"
	"math"
)

//1. Константы - это неизменяемые переменные, которые служат для:
//	1) Более строгого понимания кода
//	2) Для того, чтобы случайно не поменять значение (предполагается, что значение константы не изменно)
//	3) Для удобных преобразований

const (
	MAIN_PORT = "8001"
)

func main() {
	//2. Объявление одной константы
	const a = 10
	fmt.Println(a) //10
	//3. Объявление блока констант с областью видимости внутри функции main
	const (
		ipAddress = "127.127.00.03"
		port      = "8000"
		dbName    = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress) //ipAddress value: 127.127.00.03
	fmt.Println(checkPortIsValid())            //true

	//4. Константу никак нельзя поменять в ходе работы программы
	// const b = 200
	// b = 30

	//5. Значения констант ДОЛЖНЫ БЫТЬ ИЗВЕСТНЫ на момент компиляции
	var sqrt = math.Sqrt(25)
	// const sqrt = math.Sqrt(25) //Нельзя присвоить в константу что-либо, что является результатом вызова функции, метода -> Lec14\main.go:36:15: math.Sqrt(25) (value of type float64) is not constant
	fmt.Println("Var sqrt:", sqrt) //Var sqrt: 5

	//6. Типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - повышение читабельности кода

	//7. Нетипизированные константы и их профит
	// const NUMERIC = 100
	// var numInt = NUMERIC
	// fmt.Printf("Value %v and Type %T\n", numInt, numInt) //Value 100 and Type int
	//При использовании нетипизированных констант, мы разрешаем компилятору использовать неявное
	// приведение типов в момент присваивания значений констант в переменные
	// const NUMERIC = 100
	const NUMERIC = 10
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC
	fmt.Printf("numInt8 value %v type %T\n", numInt8, numInt8)       //numInt8 value 100 type int8
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC) //100 + 100 is -56; 10 + 10 is 20
	fmt.Printf("numInt8 value %v type %T\n", numInt32, numInt32)     //numInt8 value 100 type int32
	fmt.Printf("numInt8 value %v type %T\n", numInt64, numInt64)     //numInt8 value 100 type int64
	fmt.Printf("numInt8 value %v type %T\n", numFloat32, numFloat32) //numInt8 value 100 type float32
	fmt.Printf("numInt8 value %v type %T\n", numComplex, numComplex) //numInt8 value (100+0i) type complex64

	//8. Константы в Go зашиваются в момент компиляции в RUNTIME программы и не выбрасываются до её окончания
}

func checkPortIsValid() bool {
	if MAIN_PORT == "8001" { //если не 8001, то выдает false вместо true
		return true
	}
	return false
}
