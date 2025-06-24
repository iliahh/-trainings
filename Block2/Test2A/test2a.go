package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var logo string
	fmt.Scan(&logo)

	X := (utf8.RuneCountInString(logo) * 23) / 100
	Y := (utf8.RuneCountInString(logo) * 23) % 100

	if X > 0 {
		fmt.Printf("%d р. %d коп.\n", X, Y)
	} else {
		fmt.Printf("%d коп.\n", Y)
	}
}
