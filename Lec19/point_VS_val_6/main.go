package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	//Значение исходное
	rectangleAsValue := Rectangle{10, 10}
	//Ссылка на исходное значение
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call Area func for &rectangle:", Area(rectangleAsPointer))    //Call Area func for &rectangle: 100
	fmt.Println("Call Area method for &rectangle:", rectangleAsPointer.Area()) //Call Area method for &rectangle: 100

	//1. Вызываем метод у value - исходного значения
	fmt.Println("Call Area method for rectangle:", rectangleAsValue.Area()) //Call Area method for rectangle: 100

	//2. Вызываем функцию с параметром value - исходное значение
	//Area(rectangleAsValue) - ошибка компилятора

	//3. Распечатаем исходный прямоугольник
	fmt.Println("Before changing width:", rectangleAsValue) //Before changing width: {10 10}

	//4. Вызываю метод SetWidth у &rectangle
	rectangleAsPointer.SetWidth(999)
	fmt.Println("After change via method SetWidth for &rectangle:", rectangleAsValue) //After change via method SetWidth for &rectangle: {10 999}

	//5. Вызов метода SetWidth у rectangle
	rectangleAsValue.SetWidth(888)
	fmt.Println("After hange via method SetWidth for rectangle:", rectangleAsValue) //After hange via method SetWidth for rectangle: {10 888}
}
