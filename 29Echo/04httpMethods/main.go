package main

import (
	"fmt"
	"net/http"
)

func mylogin(w http.ResponseWriter, r *http.Request) {

	// THAT IS HOW SAME ENDPOINT CAN BEHAVE DIFFERETLY IN DIFFERNET HTTP METHOD
	if r.Method == "GET" {
		fmt.Fprintln(w, "using GET for login endpoint")
	}
	if r.Method == "POST" {
		fmt.Fprintln(w, "using POST for login endpoint")
	}
	if r.Method == "PUT" {
		fmt.Fprintln(w, "using PUT for login endpoint")
	}
	fmt.Fprintf(w, "on login page")
}


func main() {

	http.HandleFunc("/login/", mylogin) 

	fmt.Println("listening on port 3000.....") 
	http.ListenAndServe("localhost:3000", nil) 
}
