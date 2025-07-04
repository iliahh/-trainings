package main

import (
	"fmt"
	"math"
)

//3. В чем преимущество методов над функциями?
// Во-первых, наличие методов улучшает "консистентность" кода, т.к. напрямую влияет на его понимание.
// Во-вторых, в Go запрещается создавать функции с одинаковыми названиями, в то время как методы для
// разных структур с одинаковыми названиями разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// 4. Создадим функцию и метод Perimeter для структуры Circle
func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

// 5. Создадим функцию и метод Perimeter для структуры Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// 6. В Go разрешено создавать методы с одинаковыми именами в пределах одной структуры.
// Главное, чтобы получатель метода в разных структурах (где реализован метод со схожим именем) отличался.
func PerimeterRectangle(r Rectangle) int { //
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", Perimeter(c)) //Call function: 65.97344572538566
	fmt.Println("Call method:", c.Perimeter())  //Call method: 65.97344572538566

	r := Rectangle{10, 20}
	fmt.Println("Call function for rectangle:", PerimeterRectangle(r)) //Call function for rectangle: 60
	fmt.Println("Call method for rectangle:", r.Perimeter())           //Call method for rectangle: 60
}
