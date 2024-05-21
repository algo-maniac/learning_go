package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/algo_maniac/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")

	r := router.Router()
	fmt.Println("Server is running.....")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000.....")
}
