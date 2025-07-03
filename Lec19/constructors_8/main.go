package main

import "fmt"

//1. Создадим структуру Rectangle
type Rectangle struct {
	length, width int
}

//2. Создадим конструктор для Rectangle
// Для имён конструкторов в Go договорились использовать функцию с следующим названием:
// * New() - если данный конструктор на файл один (в файле присутствует описание только одной структуры)
// * New<StructName>() - если в файле присутствуют ещё какие-то структуры

//3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него
func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

//4. Добавим 2 метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)
	// fmt.Println(r) //{10 20}
	fmt.Printf("Type as %T and value %v\n", r, r) //Type as *main.Rectangle and value &{10 20}
	fmt.Println("Perimeter:", r.Perimeter())      //Perimeter: 60
	fmt.Println("Area:", r.Area())                //Area: 200
}

type Circle struct {
	radius float64
}

//!!! Если мы хотим создать двум структурам два конструктора, то мы дополняем название конструктора
// New() еще дполнительными определениями, к примеру NewCircle() и NewRectangle()
func NewCircle(newRadius float64) *Circle { //func New(newRadius float64) *Circle {
	return &Circle{newRadius}
}
