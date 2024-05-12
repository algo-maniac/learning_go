package main

import "fmt"

func main() {
	age := 20
	hasCar := false

	// Open braces should be in the same line as the condition. This is the only syntax for the lexer
	if age > 18 && hasCar {
		fmt.Println("Adult")
	} else {
		fmt.Println("Underage")
	}
}
