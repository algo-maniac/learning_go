package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Taking inputs and converting into other datatypes
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the size of your shirt:")
	input, _ := reader.ReadString('\n')
	newSize, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The new size of the shirt: ", newSize+1)
	}

}
