package main

import "fmt"

func main() {
	var ptr *int
	number := 24
	fmt.Println("Value of default pointer : ", ptr)

	// We can use normal assignment operators in case of pointer
	ptr = &number
	fmt.Println("Value of integer pointer : ", *ptr)
	fmt.Println("Address of integer pointer : ", ptr)

}
