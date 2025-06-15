package main

import (
	"fmt"
	"strings"
)

func main() {
	var login, mail string
	fmt.Scan(&login, &mail)

	if len(login) < 10 || strings.Contains(login, "@") {
		fmt.Println("Некорректный логин")
		return
	} else if !strings.Contains(mail, "@") || !strings.Contains(mail, ".") {
		fmt.Println("Некорректная почта")
		return
	} else {
		fmt.Println("ОК")
	}
}
