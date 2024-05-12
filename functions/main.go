package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to functions")

	greaterValue := findMaximum(10, 30)
	fmt.Println(greaterValue)

	list := []int{2, 3, 4, 5}
	totalSum := proAdder(list)
	fmt.Println("Total sum is : ", totalSum)

	// You can return more than 1 values from the string
	statusCode, password := passwordChecker("Soumyajit@13*")
	fmt.Println(statusCode, password)

}

func findMaximum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func proAdder(values []int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func passwordChecker(password string) (bool, string) {
	if strings.Contains(password, "@") && strings.Contains(password, "*") {
		return true, password
	} else {
		return false, password
	}
}
