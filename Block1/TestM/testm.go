package main

import (
	"fmt"
)

func main() {
	var first, second, third string
	fmt.Scan(&first, &second, &third)

	if (first == "раз" && second == "два" && third == "три") ||
		(first == "один" && second == "два" && third == "три") ||
		(first == "1" && second == "2" && third == "3") {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}

	/*
		if strings.Contains(raz, "раз") && strings.Contains(dva, "два") && strings.Contains(tri, "три") {
			fmt.Println("ОК")
		} else if strings.Contains(raz, "один") && strings.Contains(dva, "два") && strings.Contains(tri, "три") {
			fmt.Println("ОК")
		} else if strings.Contains(raz, "1") && strings.Contains(dva, "2") && strings.Contains(tri, "3") {
			fmt.Println("ОК")
		} else {
			fmt.Println("НЕ ПРАВИЛЬНО")
		}
	*/

}
