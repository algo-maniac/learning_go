package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myUrl string = "https://dummyjson.com/products/search?q=Laptop"
const postUrl string = "https://dummyjson.com/products/add"

var getResponseString string

func GetRequests() {
	fmt.Println("LCO web request")
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type : %T\n", response)
	defer response.Body.Close() // callers responsibility to close response
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	getResponseString = bodyString
	fmt.Println(bodyString)

}

func PostRequests() {

	data := url.Values{}

	data.Add("name", "Soumyajit")
	data.Add("college", "Jadavpur University")
	data.Add("email", "soumya@gmail.com")
	fmt.Printf("Type of data is : %T", data)
	requestBody := strings.NewReader(
		`{
			"user" : "Soumyajit",
			"college": "Jadavpur University"
		}
		`)

	response, err := http.Post(postUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}

func ParseUrl(myUrl string) {
	// Parsing urls

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Host, result.Scheme, result.Path, result.RawQuery, result.Port())
	qparams := result.Query()

	for key, val := range qparams {
		fmt.Println(key, val)
	}
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // Remove this field
	Tags     []string `json:"tags,omitempty"` // Remove this field if it is empty for some case
}

func EncodeJson() {
	myCourses := []course{
		{
			"ReactJS Course",
			500,
			"Udemy",
			"soumya@gmail.com",
			[]string{"web-dev", "js", "frontend"},
		},
		{
			"FullStack Course",
			500,
			"Udemy",
			"soumya@gmail.com",
			[]string{"web-dev", "js", "fullstack"},
		},
		{
			"Python Course",
			500,
			"Udemy",
			"soumya@gmail.com",
			[]string{},
		},
	}

	// Package this data as json

	finalJsonByte, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Json is %v\n", string(finalJsonByte))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename":"React JS Bootcamp",
			"price":500,
			"website":"Udemy.com",
			"tags":["web-dev", "js"],
			"more_tags":{
				"key1":30,
				"key2":50
			}
		}
		`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// Another method

	var myonlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for key, value := range myonlineData {
		fmt.Printf("key is %v, value is %v\n", key, value)
	}

}
func main() {
	// ParseUrl(myUrl)
	// GetRequests()
	// PostRequests()
	// EncodeJson()
	DecodeJson()
}
