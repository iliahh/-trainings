package main

import "fmt"

func main() {
	var name, surname string
	var age int
	fmt.Scan(&name, &surname, &age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS\n", name, surname, age)
}
