package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"` //FullName jab json me use hoga tab fullname ho jayega
	Website  string `json:"website"`
}

// fake db
var courses []Course

//middleware,helper-file

func (c *Course) isEmpty() bool {
	return (c.CourseName == "" && c.CourseId == "")
}

func main() {

}

//controller - file
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Api by learnCodeOnline </h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // this is how we set the header, responseWriter (w) me Header set karne ke leye Header diya hua and >set se set kardo, and set method key-value pair leti hai
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r) // r ho ja wo reader hai mtlb url hai jaha se read karna hai mux ke pass ek method hai Vars jo key,value jo ki ? iske baad hote hai wo dedegi, ha url ka bhi use kar sakte hai

	// loop through courses,find matching id and return the response

	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("no course with given id");
}
