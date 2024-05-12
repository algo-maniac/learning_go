package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	working_days := []string{"Mon", "Tue", "Wed", "Thru", "Fri"}
	fmt.Printf("The working days are : %+v", working_days)

	// for i := 0; i < len(working_days); i++ {
	// 	fmt.Println(working_days[i])
	// }

	// Same loop

	for i := range working_days {
		fmt.Println(working_days[i])
	}

	for index, day := range working_days {
		if index == 2 {
			goto jmp
		}
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	// goto statements

jmp:
	fmt.Println("Hey I jumped here from the if condition inside for")
}
