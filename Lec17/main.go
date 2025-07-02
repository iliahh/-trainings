package main

import "fmt"

//1. Структура - это заименованный набор полей (свойств), определяющий новый тип данных.

//2. Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

//3. Если имеется ряд состояний одного типа, можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, coutrseNumber             int
}

//11. Структура с анонимными полями
type Human struct {
	//Два явных состояния
	firstName string
	lastName  string
	//Три без каких-либо имен
	string
	int
	bool
}

func PrintStudent(std Student) {
	fmt.Println("===================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {

	//4. Создание представителей структуры
	stud1 := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		age:       21,
	}
	// fmt.Println("Student 1:", stud1) //Student 1: {Fedya Petrov 21}
	PrintStudent(stud1) //->
	// ===================
	// FirstName: Fedya
	// LastName: Petrov
	// Age: 21
	stud2 := Student{"Petya", "Ivanov", 19} //Порядок указания свойств - такой же как в структуре
	PrintStudent(stud2)                     //->
	// ===================
	// FirstName: Petya
	// LastName: Ivanov
	// Age: 19

	//5. Что если не все поля структуры определить?
	stud3 := Student{
		firstName: "Vasya",
	}
	PrintStudent(stud3) //->
	// FirstName: Vasya
	// LastName:
	// Age: 0

	//6. Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupID       int
		proffesorName string
	}{
		age:           23,
		groupID:       2,
		proffesorName: "Alexeev",
	}
	fmt.Println("AnonStudent:", anonStudent) //AnonStudent: {23 2 Alexeev}

	//7. Доступ к состояниям
	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("firstName:", studVova.firstName) //firstName: Vova
	fmt.Println("lastName:", studVova.lastName)   //lastName: Ivanov
	fmt.Println("age:", studVova.age)             //age: 19
	studVova.age += 2
	fmt.Println("new age:", studVova.age) //new age: 21

	//8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1) //->
	// ===================
	// FirstName:
	// LastName:
	// Age: 0
	PrintStudent(emptyStudent2) //->
	// ===================
	// FirstName:
	// LastName:
	// Age: 0

	//9. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer:", studPointer) //Value studPointer: &{Igor Sidorov 22}
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value afterPointerModify:", studPointer) //Value afterPointerModify: &{Igor Sidorov 42}
	studPointerNew := new(Student)
	fmt.Println(studPointerNew) //&{  0}

	//10. Работа с доступом к полям структур через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age) //Age via (*...).age: 42
	fmt.Println("Age via .age:", studPointer.age)          //Age via .age: 42 //Неявно происходит разыменование указателя studPointer и запрос соответствующего поля

	//12. Создание экземпляра с анонимными полями структуры
	human := &Human{
		firstName: "Bob",
		lastName:  "Johnson",
		string:    "Additional Info",
		int:       -1,
		bool:      true,
	}
	fmt.Println(human) //&{Bob Johnson Additional Info -1 true}
	//Как обратиться к анонимным:
	fmt.Println("Anon field string:", human.string) //Anon field string: Additional Info - обратились к строке
}
