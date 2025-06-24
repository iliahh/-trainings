1) Первая версия:
package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)

	if w <= 100 && w >= 1 {
		if w%2 == 0 && w > 2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}

2) Вторая версия:
package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)

	if w%2 == 0 && w > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}