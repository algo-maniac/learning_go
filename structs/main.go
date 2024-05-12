package main

import "fmt"

func main() {
	soumya := User{"Soumyajit", "soumya@gmail.com", 20, true}
	fmt.Printf("User details are %+v", soumya)
	fmt.Println(soumya.Name)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
