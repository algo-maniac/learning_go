package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{}

func callApi(api string) {

	defer wg.Done()
	res, err := http.Get(api)

	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, api)

		mut.Unlock()
		fmt.Printf("Response returned from %s with status code %s\n", api, res.Status)
	}
}
func main() {
	apiList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://billpal.vercel.app",
		"https://www.github.com",
		"https://codeforces.com/",
		"https://leetcode.com/",
		"https://go.dev/",
	}
	for _, val := range apiList {
		go callApi(val)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}
