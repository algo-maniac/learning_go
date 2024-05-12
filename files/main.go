package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fileContent := "Hey, I am Soumyajit and I am learning Go lang"

	file, err := os.Create("./file.txt")

	if err != nil {
		// fmt.Println("Some error occured during creation of file")
		panic(err)
	} else {
		fmt.Println("File created : ", file)
	}
	// fmt.Println(fileContent)

	length, err := io.WriteString(file, fileContent)
	fmt.Println("Length of file is : ", length)

	filepath, err := filepath.Abs("./file.txt")
	fmt.Println(filepath)
	readFile(filepath)
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Text inside the file is : ", string(databyte))
	}

}
