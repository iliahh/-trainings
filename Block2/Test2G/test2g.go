package main

import (
	"fmt"
)

func main() {
	var pulse, min, max float64
	var count int
	first := true
	for {
		fmt.Scan(&pulse)
		if pulse < 0 {
			break
		}
		if pulse >= 100 && pulse <= 140 {
			count++
			if first {
				min, max = pulse, pulse
				first = false
			} else {
				if pulse < min {
					min = pulse
				}
				if pulse > max {
					max = pulse
				}
			}
		}
	}
	fmt.Println(count)
	if count > 0 {
		fmt.Printf("%.1f %.1f\n", min, max)
	}
}
