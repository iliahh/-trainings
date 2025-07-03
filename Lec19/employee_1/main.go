package main

import "fmt"

//Тема - методы, интерфейсы, способы поддержки объектно-ориентированного программирования

//Go - не является объектно-ориентированный язык, не поддерживает классы, тут их нет в принципе,
// но сторонняя поддержка ООП здесь тем не менее присутствует,
// но выглядит не совсем классическим способом.

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

//1. Методы - функции, привязанные к определённым структурам
//func (s Struct) MethodName(parameters type) type {}
//		Reciever - получатель метода
func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "Bob",
		position: "Middle Golang developer",
		salary:   3000,
		currency: "USD",
	}

	//2. Вызов метода
	emp.DisplayInfo() //->
	// Name: Bob
	// Position: Middle Golang developer
	// Salary: 3000 USD
}
