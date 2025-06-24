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
