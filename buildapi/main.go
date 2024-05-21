package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Response messages

type ControllerResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Constructor function to create a Success instance with default values
func NewControllerResponse(statusCode int, message string) ControllerResponse {
	return ControllerResponse{
		Status:  statusCode, // Default status code
		Message: message,    // Default message
	}
}

// Database

var courses []Course

// Middlewares, helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Soumyajit</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(courses)

	if len(courses) == 0 {
		response := NewControllerResponse(400, "No courses found")
		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode(courses)
	}

}

// get a course

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// id param for the request
	params := mux.Vars(r)

	// loop through all the courses and match the id

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	response := NewControllerResponse(400, `No courses found with given id`)
	json.NewEncoder(w).Encode(response)

}

// create a course

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty

	if r.Body == nil {
		response := NewControllerResponse(400, `Please send some data`)
		json.NewEncoder(w).Encode(response)
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		response := NewControllerResponse(400, `No data inside json`)
		json.NewEncoder(w).Encode(response)
		return
	}

	// generate unique id
	// append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

// update a course

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// id param for the request
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	response := NewControllerResponse(400, `Id not matched`)
	json.NewEncoder(w).Encode(response)

}

// delete a course

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// id param for the request
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	response := NewControllerResponse(400, `Id not matched`)
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Hey Welcome to my api")

	r := mux.NewRouter()

	// seeding

	myCourses := []Course{
		{
			"001",
			"ReactJS Course",
			500,
			&Author{"Soumyajit Naskar", "billpal.vercel.app"},
		},
		{
			"002",
			"Python Course",
			600,
			&Author{"Tuhin Saha", "thinktanktrivia.vercel.app"},
		},
	}

	fmt.Println(myCourses)
	courses = myCourses

	// routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}
