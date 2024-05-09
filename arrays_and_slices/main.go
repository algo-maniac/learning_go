package main

import (
	"fmt"
	"sort"
)

func main() {

	// Normal Arrays

	// They are of fixed sizes the values inside the array can be changed but the length cannot be changed
	var array [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(array)

	// Redeclaring the array
	array = [4]int{2, 3, 4, 6}
	fmt.Println(array)

	mySlice := []int{2, 3, 4, 5}
	fmt.Println(mySlice)

	// Inserting element into the slice
	mySlice = append(mySlice, 7, 1, 5)
	fmt.Println(mySlice)

	// Deleting element at index 2
	index := 2
	mySlice = append(mySlice[:index], mySlice[index+1:]...)
	fmt.Println(mySlice)

	// Sorting elements in the slice
	sort.Ints(mySlice)
	fmt.Println(mySlice)

}
