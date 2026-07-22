package main

import ( 
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	scanner.Scan()
	age := scanner.Text()
	fmt.Printf("Hi, %s! You are %s years old.", name, age)
}
