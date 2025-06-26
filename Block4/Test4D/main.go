package main

import "fmt"

func main() {
	var money, heads int
	fmt.Scan(&money, &heads)

	for bull := 1; bull <= money; bull++ { //Перебор циклом количества быков (1,2...n) (требуется хотя бы один бык)
		for cow := 0; cow <= heads-bull; cow++ { //Перебор коров при каждом числе быков (но коровы могут отсутствовать)
			calve := heads - bull - cow //Телята - остаток, их может и не быть
			cost := bull*20 + cow*10 + calve*5
			if cost == money {
				fmt.Printf("%d %d %d\n", bull, cow, calve)
			}
		}
	}
}
