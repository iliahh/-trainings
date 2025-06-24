// package main

// import "fmt"

//	func main() {
//		var word1, word2, word3, word4, author string
//		fmt.Scan(&word1, &word2, &word3, &word4, &author)
//		fmt.Printf("%s - "+author+"\n", word4)
//		fmt.Printf("%s - "+author+"\n", word3)
//		fmt.Printf("%s - "+author+"\n", word2)
//		fmt.Printf("%s - "+author+"\n", word1)
//	}
package main

import "fmt"

func main() {
	var messages [4]string
	var author string

	// Считываем 4 строки
	for i := 0; i < 4; i++ {
		fmt.Scan(&messages[i])
	}

	// Считываем имя автора
	fmt.Scan(&author)

	// Выводим в обратном порядке
	for i := 3; i >= 0; i-- {
		fmt.Printf("%s - %s\n", messages[i], author)
	}
}
