package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World!!")

	// We can declare variable by the following methods

	var nameOne string = "Soumya"
	var marks int = 75
	var nameThree bool // default value if value not mentioned
	name := "hey"
	var score int64 = 100000 // for higher bit size

	fmt.Println(name, score) // prints a \n after this line
	fmt.Println(nameOne, marks, nameThree)
	fmt.Print("Hello\n") // prints only the given stuff
	fmt.Printf("My name is %v, and my score is %v\n", nameOne, marks)

	// Here we can also save output statements inside variables

	// Also general thing to note is that just declare variables like var variableName = "anything"/anything because go will infer it

	var savedString = fmt.Sprintf("My name is %v, and my score is %v", nameOne, marks)
	fmt.Println(savedString)

	// Declaring an array
	var arr = [3]int{2, 3, 5}
	names := [3]string{"Atanu", "Soumyajit", "Tuhin"}
	names[1] = "Arith"
	fmt.Println(arr, len(arr))
	fmt.Println(names, len(names))

	// Declaring a slice it is variable len in nature
	scores := []int32{2, 3, 4}
	fmt.Println(scores)
	// Some slice functions
	scores = append(scores, 5) // overwriting
	fmt.Println(scores, len(scores))
	sliceRange := scores[1 : len(scores)-1]
	sliceRangeTwo := scores[:3]
	sliceRangeThree := scores[:]
	fmt.Println(sliceRange, sliceRangeTwo, sliceRangeThree)

}
