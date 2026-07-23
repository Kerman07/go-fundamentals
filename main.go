package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	sum, i := 0, 1
	for i <= N {
		sum += i
		i += 1
	}
	fmt.Println(sum)
}
