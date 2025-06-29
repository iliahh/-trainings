package main

import "fmt"

//1. Указатели - переменная, хранящая в качестве значения адрес в памяти другой переменной

func main() {
	//2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&.... - операция взятия адреса в памяти
	//Выше у нас создан указатель на переменную variable
	//В pointer лежит 18293xcd000132 - это место в памяти, где хранится int значение 30
	fmt.Printf("Type of pointer %T\n", pointer)  //Type of pointer *int
	fmt.Printf("Value of pointer %v\n", pointer) //Value of pointer 0xc0000940a8

	//3. А какое нулевое значение для указателя?
	var zeroPointer *int                                           //zeroValue имеет значение nil (это указатель в никуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer) //Type *int and value <nil>
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initialization type %T and value %v\n", zeroPointer, zeroPointer) //After initialization type *int and value 0xc00000a0d8
	}

	//4. Разыменование указателя (получение значения): *pointer - возвращает значение, хранимое по адресу
	// zeroPointerToPointer := &zeroPointer
	// fmt.Println(zeroPointerToPointer) //0xc000072068 (так лучше не делать, синтаксический анализатор будет ругаться, и такого в коде в целом не будет дальше)

	var numericValue int = 32
	var pointerToNumeric *int = &numericValue

	fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)
	//Value in numericValue is 32
	//Address is 0xc0000940e0

	//5. Создание указателей на нулевые значения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	zeroPoint := new(int) //Создает под капотом zeroValue для int, и возвращает адрес, где этот 0 хранится

	fmt.Printf("Value in *zeroPoint %v\nAddress is %v\n", *zeroPoint, zeroPoint)
	//Value in *zeroPoint 0
	//Address is 0xc00000a118

	//6. Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	// fmt.Println("Value in zeroPointerToInt is:", *zeroPointerToInt) //Value in zeroPointerToInt is: 0
	fmt.Printf("Address is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt) //Address is 0xc00000a120 and Value in zeroPointerToInt is 0
	*zeroPointerToInt += 40
	fmt.Printf("Address is %v and New Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt) //Address is 0xc00000a120 and New Value in zeroPointerToInt is 40

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100      //fmt.Println(b) //446
	fmt.Println(b) //346

	//7. Указательная арифметика ОТСУТСТВУЕТ ПОЛНОСТЬЮ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес продвинуться в другие ячейки
	// 	cpp_like := "cpp"
	// 	cpp_pointer := &cpp_like
	// 	cpp_pointer++

	//8. Передача указателей в функции
	// Колоссальный прирост производительности за счёт того, что передаётся не значение (которое должно копироваться),
	// а передаётся лишь адрес в памяти, за которым уже хзранится какое-то значение
	sample := 1
	// samplePointer := &sample

	fmt.Println("Origin value of sample:", sample) //Origin value of sample: 1
	// changeParam(samplePointer)
	changeParam(&sample)
	// changeWOPointer(sample)
	fmt.Println("After changing sample is:", sample) //After changing sample is: 101

	//9. Возврат поинтера из функции (в C++ результат работы такого механизма - не определён)
	ptr1 := returnPointer() //На каждый вызов этой функции создается новый нумерик и кладет его адрес
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1) //Ptr1: *int and address 0xc00000a138 and value 321
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr2, ptr2, *ptr2) //Ptr1: *int and address 0xc00000a140 and value 321
}

//9.1 Инициализация функции, возвращающей указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric //В момент возврата Go перемещает данную переменную в кучу
}

// func changeWOPointer(val int) {
// 	val += 100
// }

//8.1 Определение функции, принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}
