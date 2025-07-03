package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

// 1. Метод, в котором получатель копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

// 2. Метод, в котором получатель передаётся по ссылке (в теле метода работаем с ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

//4. Используйте методы с PointerReciever'ом в ситуациях, когда:
//	1) Изменения в работе метода над экземпляром должны быть видны на вызывающей стороне
//	2) когда экземпляр достаточно "увесистый", то есть копировать его дорого с точки зрения расходов ресурсов
//	3) С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting parameters:", e) //Before setting parameters: {Alex 3000}
	e.SetName("Bob")
	fmt.Println("After SetName call:", e) //After SetName call: {Alex 3000}
	e.SetSalary(4500)                     //3. Вызов метода у ссылки на сотрудника
	//5. Аналогично явному вызову (&e).SetSalary(9999)
	fmt.Println("After SetSelary call:", e) //After SetSelary call: {Alex 4500}
}
