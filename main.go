package main

import (
	"fmt"
)

func square(n int) int {
	return n * n
}

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(square(N))
}
