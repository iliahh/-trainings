package main

import (
	"fmt"
)

func main() {
	var value int
	for {
		fmt.Scan(&value)
		if value > 0 {
			fmt.Println(value)
			continue
		} else {
			return
		}
	}
}
