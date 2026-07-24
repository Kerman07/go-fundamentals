package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convert(text string)  {
	a, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("bad")
		return
	}
	fmt.Printf("ok %d", a)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	convert(text)
}
