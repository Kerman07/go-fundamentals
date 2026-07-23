package main

import (
	"fmt"
)

func main() {
	var integer int
	fmt.Scan(&integer)

	switch {
	case integer%15 == 0:
		fmt.Println("FizzBuzz")
	case integer%5 == 0:
		fmt.Println("Buzz")
	case integer%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(integer)
	}
}
