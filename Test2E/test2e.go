package main

import (
	"fmt"
	"strings"
)

func main() {
	var password1, password2 string //LOLKEK9000ы

	for {
		fmt.Scan(&password1, &password2)
		if len(password1) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		} else if (len(password1) >= 8 && strings.Contains(password1, "123")) ||
			(len(password1) >= 8 && strings.Contains(password1, "qwe")) {
			fmt.Println("Слишком простой пароль!")
			continue
		} else if password1 != password2 {
			fmt.Println("Введенные пароли различаются!")
			continue
		} else {
			fmt.Println("Ну наконец-то!")
			return
		}
	}
	//Дело в том, что сначала считывались оба пароля один раз (считывание стояло вне цикла),
	// а потом запускался бесконечный цикл, который повторно не запрашивал ввод —
	// то есть, данные не обновлялись, и цикл не выполнял повторные попытки ввода.
}
