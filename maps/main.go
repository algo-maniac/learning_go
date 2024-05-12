package main

import "fmt"

func main() {

	// syntax : make(map[<key>]<value>)
	languages := make(map[string]string)
	languages["PY"] = "Python"
	languages["CPP"] = "C++"
	languages["JV"] = "Java"
	fmt.Println(languages)
	fmt.Println(languages["PY"])

	// Deleting element
	delete(languages, "JV")
	fmt.Println(languages)

	// Looping through the map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
