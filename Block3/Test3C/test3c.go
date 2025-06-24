package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаем сканер для чтения строк из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	var prev string // Предыдущее слово

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text()) // Текущее слово, обрезаем пробелы

		// Если это не первое слово, проверяем правило
		if prev != "" {
			last := []rune(prev)[len([]rune(prev))-1] // последняя буква предыдущего слова
			first := []rune(word)[0]                  // первая буква текущего слова

			if first != last {
				// Нарушение правила — выводим слово и завершаем
				fmt.Println(word)
				return
			}
		}

		prev = word // Запоминаем текущее слово как предыдущее
	}

	// Если достигли конца ввода и все слова корректны, ничего не выводим
}
