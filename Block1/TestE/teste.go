package main

import "fmt"

func main() {
	var drink, meal, subMeal, time string
	fmt.Scan(&drink, &meal, &subMeal, &time)
	//fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'\n", drink, meal, subMeal, time)
	fmt.Printf("I wanna some '" + drink + "', my favorite meal - '" + meal + "'. Give me also '" + subMeal + "'. I will come soon... '" + time + "'\n")
}
