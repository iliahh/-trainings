package main

import (
	"fmt"
	"strings"
)

func main() {
	var stroka string
	fmt.Scan(&stroka)

	if strings.Contains(stroka, "чат") {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
