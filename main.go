package main

import "fmt"

func main() {
	// fmt.Println("Hello World!!")

	// We can declare variable by the following methods

	var nameOne string = "Hello"
	var marks int = 75
	var nameThree bool // default value if value not mentioned
	name := "hey"
	var score int64 = 100000 // for higher bit size

	fmt.Println(name, score)
	fmt.Println(nameOne, marks, nameThree)
}
