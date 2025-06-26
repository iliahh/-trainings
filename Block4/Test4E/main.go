package main

import (
	"fmt"
)

func main() {
	var m, n int //m - ученики на С++, n - ученики на Rust
	fmt.Scan(&m, &n)

	cplus := make([]string, m)
	rust := make([]string, n)

	for i := 0; i < m; i++ {
		fmt.Scan(&cplus[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&rust[i])
	}

	cMap := make(map[string]bool)
	rMap := make(map[string]bool)

	for _, name := range cplus {
		cMap[name] = true
	}

	for _, name := range rust {
		rMap[name] = true
	}

	count := 0
	for name := range cMap {
		if !rMap[name] {
			count++
		}
	}

	for name := range rMap {
		if !cMap[name] {
			count++
		}
	}

	fmt.Println(count)
}
