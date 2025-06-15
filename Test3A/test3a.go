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

	for i := range m {
		scanner.Scan()
		wantedSeries[i] = scanner.Text()
	}

	//В общем, тут переменные wanted и service - это переменные для сравнения без учета регистра,
	//они обе поочередно берут значения из слайсов serviceSeries и wantedSeries
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
