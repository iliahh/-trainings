package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Сканер для чтения строк из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	var prev string // Предыдущее слово

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text()) // Текущее слово

		if prev != "" {
			prevRunes := []rune(prev)
			wordRunes := []rune(word)

			// Обработка мягкого знака в конце предыдущего слова
			lastIndex := len(prevRunes) - 1
			lastChar := prevRunes[lastIndex]
			if (lastChar == 'ь' || lastChar == 'Ь') && len(prevRunes) >= 2 {
				lastChar = prevRunes[lastIndex-1]
			}

			firstChar := wordRunes[0]

			if firstChar != lastChar {
				fmt.Println(word)
				return
			}
		}

		prev = word
	}
}
