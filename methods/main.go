package main

import "fmt"

func main() {
	soumya := User{"Soumyajit", "soumya@gmail.com", 20, true}
	soumya.getUser()
	soumya.GetStatus()
	soumya.newMail()
	soumya.getUser()

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

// Use pointers to actually manipulate the variables rather than just pass by value
func (u *User) newMail() {
	u.Email = "soumya.dev@gmail.com"
}

func (u User) getUser() {
	fmt.Printf("User details are %+v", u)
}
