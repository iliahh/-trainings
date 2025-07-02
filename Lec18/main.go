package main

import "fmt"

//1. Вложенные структуры (вложение структур). Это использование одной структуры,
// как тип поля в другой структуре
type University struct {
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

//4. Встроенные структуры (когда мы добавляем поля одной структуры к другой
type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	//8.1 Вложим в профессора какое-то поле, имеющей некий несравнимый типаж - вылезает ошибка компиляции
	//papers []string или papers map[string]string (это нельзя сравнить, только с nil можно)

	University //В этом месте происходит добавление всех полей структуры Univercity в Professor
}

func main() {
	//2. Создание эксземпляров структур с вложением
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{ // * А вто здесь обращение идёт к полю юниверсити как к вложенному (То есть есть структура Student с полем в ней univercity, которая инициализируется)
			yearBased: 1991,
			infoShort: "cool University",
			infoLong:  "very cool University",
		},
	}

	//3. Получение доступа к вложенным полям структур
	fmt.Println("firstName:", stud.firstName)                 //firstName: Fedya
	fmt.Println("lastName:", stud.lastName)                   //lastName: Petrov
	fmt.Println("Year based Uni:", stud.university.yearBased) //Year based Uni: 1991
	fmt.Println("longInfo:", stud.university.infoLong)        //longInfo: very cool University

	//5. Создание экземпляра со встраиванием структур
	prof := Professor{
		firstName: "Anatoly",
		lastName:  "Smirnov",
		age:       125,
		greatWork: "Ultimate C programming",
		University: University{ // * Здесь обращение к юниверсити происходит как к анонимному, будто тут только тип указан
			yearBased: 1734,
			infoShort: "short Info",
			//age:       2021 - 1734,
		},
	}

	//6. Обращение к состояниям со встроенной структурой
	fmt.Println("firstName:", prof.firstName)  //firstName: Anatoly
	fmt.Println("Year based:", prof.yearBased) //Year based: 1734
	fmt.Println("Info Short:", prof.infoShort) //Info Short: short Info
	//fmt.Println("Age:", prof.University.age)   //prof.age - получим доступ к полю ВЫШЕЛЕЖАЩЕЙ СТРУКТУРЫ

	//7. Сравнение экземпляров ==
	//При сравнении экземпляров происходит сравнение их всех полей друг с другом
	profLeft := Professor{
		firstName: "Victor",
	}
	profRight := Professor{}

	fmt.Println(profLeft == profRight) //true при {} == {} //false при {firstName: "Victor",} == {}

	//8. Если ХОТЯ БЫ ОДНО ИЗ ПОЛЕЙ СТРУКТУР НЕ СРАВНИМО, то и вся структура несравнима
}
