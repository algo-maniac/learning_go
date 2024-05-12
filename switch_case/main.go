package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Generate numbers from min to max range

	min := 20
	max := 100
	diceNumber := rand.Intn(max-min+1) + min
	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 60:
		fmt.Printf("Dice number is equal to 60 : %+v", diceNumber)
	default:
		fmt.Printf("Dice number is not equal to 60 : %+v", diceNumber)

	}
}
