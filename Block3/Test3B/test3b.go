package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		fmt.Println("НЕ СОГЛАСЕН")
		return
	}

	word := strings.TrimSpace(input.Text())
	runeSlice := []rune(word)

	if len(runeSlice) < 2 {
		fmt.Println("НЕ СОГЛАСЕН")
		return
	}

	firstElement := strings.ToLower(string(runeSlice[0]))
	lastElement := strings.ToLower(string(runeSlice[len(runeSlice)-1]))

	//if (firstElement == "д") && (lastElement == "а") {
	if (firstElement == "д" && lastElement == "а") || (firstElement == "а" && lastElement == "д") {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}
}
