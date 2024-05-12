package main

import "fmt"

func main() {
	// defer simply pushes that particular line of code to execute at the end of current block of function execution.

	// defer acts like a stack
	

	defer fmt.Println("World!!")
	defer fmt.Println("Universe!!")
	fmt.Println("Hello")
	fmt.Println("My")
	fmt.Println("Dear")
}
