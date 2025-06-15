1) Первая версия кода:
package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	runeSlice := []rune(word)

	for i := range runeSlice {
		if len(runeSlice) >= 2 {
			if i%2 == 0 {
				runeSlice = runeSlice[2:]
				fmt.Println(string(runeSlice))
			} else {
				runeSlice = runeSlice[:len(runeSlice)-2]
				fmt.Println(string(runeSlice))
			}
		} else {
			return
		}
	}
}

2) Вторая версия кода:
package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)

	runeSlice := []rune(word)
	start := true

	for len(runeSlice) >= 2 {
		if len(runeSlice) >= 2 {
			if start {
				runeSlice = runeSlice[2:]
			} else {
				runeSlice = runeSlice[:len(runeSlice)-2]
			}
			fmt.Println(string(runeSlice))
			start = !start
		}
	}
}

3) Третья версия кода:
package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)

	runeSlice := []rune(word)
	fmt.Println(string(runeSlice))

	start := true

	for len(runeSlice) > 2 {
		if start {
			runeSlice = runeSlice[2:]
		} else {
			runeSlice = runeSlice[:len(runeSlice)-2]
		}
		fmt.Println(string(runeSlice))
		start = !start
	}
}