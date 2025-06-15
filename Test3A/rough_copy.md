1) Первая версия кода:
package main

import (
	"fmt"
	"slices"
)

func main() {
	var n, m int
	var str1, str2 string
	fmt.Scan(&n, &m)

	var slice1 []string
	var slice2 []string

	for range n {
		fmt.Scan(&str1)
		slice1 = append(slice1, str1)
	}
	fmt.Println(slice1)

	for range m {
		fmt.Scan(&str2)
		slice2 = append(slice2, str2)
	}
	fmt.Println(slice2)

	for _, item2 := range slice2 {
		if slices.Contains(slice1, item2) {
			// fmt.Printf("Совпадение: %s\n", item2)
			fmt.Println("ЕСТЬ")
		} else {
			// fmt.Printf("Нет: %s\n", item2)
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

2) Вторая версия кода:
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var name string
	var slice1 []string
	var slice2 []string

	// Чтение сериалов на сервисе
	for i := 0; i < n; i++ {
		fmt.Scanln()   // читаем до конца строки (возможно пусто после Scan)
		getline(&name) // читаем целую строку
		slice1 = append(slice1, name)
	}

	// Чтение интересующих сериалов
	for i := 0; i < m; i++ {
		fmt.Scanln()
		getline(&name)
		slice2 = append(slice2, name)
	}

	// Проверка наличия
	for _, item := range slice2 {
		if contains(slice1, item) {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

// Функция для чтения целой строки с пробелами
func getline(dest *string) {
	var line string
	for {
		var part string
		n, _ := fmt.Scan(&part)
		if n == 0 {
			break
		}
		if line != "" {
			line += " "
		}
		line += part
		if strings.HasSuffix(part, "\n") {
			break
		}
	}
	*dest = line
}

// Проверка наличия строки в слайсе (без использования slices.Contains)
func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

3) Третья версия кода:
package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	slice1 := make([]string, n)
	slice2 := make([]string, m)

	// Считываем сериалы с сервиса
	for i := 0; i < n; i++ {
		fmt.Scan(&slice1[i])
	}

	// Считываем сериалы, которые ищет пользователь
	for i := 0; i < m; i++ {
		fmt.Scan(&slice2[i])
	}

	// Проверка наличия
	for _, wanted := range slice2 {
		found := false
		for _, available := range slice1 {
			if wanted == available {
				found = true
				break
			}
		}
		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

4) Четвёртая версия кода:
package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	slice1 := make([]string, n)
	slice2 := make([]string, m)

	for i := range n {
		fmt.Scan(&slice1[i])
	}

	for i := range m {
		fmt.Scan(&slice2[i])
	}

	// Для каждого интересующего сериала проверяем, есть ли он в сервисе
	for i := 0; i < m; i++ {
		found := false
		for j := 0; j < n; j++ {
			if slice2[i] == slice1[j] {
				found = true
				break
			}
		}
		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

5) Пятая версия кода:
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)

	// Считываем n и m
	fmt.Scan(&n, &m)

	// Создаем map для сериалов на сервисе
	serials := make(map[string]bool)

	// Считываем n названий сериалов
	scanner.Scan() // Считываем оставшийся \n после fmt.Scan
	for i := 0; i < n; i++ {
		scanner.Scan()
		title := strings.TrimSpace(scanner.Text())
		serials[title] = true
	}

	// Проверяем m интересующих сериалов
	for i := 0; i < m; i++ {
		scanner.Scan()
		want := strings.TrimSpace(scanner.Text())
		if serials[want] {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

6) Шестая версия кода:
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	serviceSeries := make([]string, n)
	wantedSeries := make([]string, m)

	for i := range n {
		scanner.Scan()
		serviceSeries[i] = scanner.Text()
	}
	fmt.Println(serviceSeries)

	for i := range m {
		scanner.Scan()
		wantedSeries[i] = scanner.Text()
	}
	fmt.Println(wantedSeries)

	for _, wanted := range wantedSeries {
		found := slices.Contains(serviceSeries, wanted)
		//for _, available := range serviceSeries {
		// 	if wanted == available {
		// 		found = true
		// 		break
		// 	}
		// }
		//Так было выше до исправления на found := slices.Contains(serviceSeries, wanted)
		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}

7) Седьмая версия кода:
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	serviceSeries := make([]string, n)
	wantedSeries := make([]string, m)

	for i := range n {
		scanner.Scan()
		serviceSeries[i] = scanner.Text()
	}
	fmt.Println(serviceSeries)

	for i := range m {
		scanner.Scan()
		wantedSeries[i] = scanner.Text()
	}
	fmt.Println(wantedSeries)

	for _, wanted := range wantedSeries {
		found := false
		for _, service := range serviceSeries {
			if strings.EqualFold(service, wanted) {
				found = true
				break
			}
		}
		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}