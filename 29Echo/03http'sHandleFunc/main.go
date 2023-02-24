package main

// gin is not install facing some issue
import (
	"fmt"
	"net/http"
)

func mylogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
	   <head>
		Login
	   <head>
	   <body>
	    <h1> Please enter your username and password </h1>
		</body>
	</html>
	`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
	   <head>
	   welcome
	   <head>
	   <body>
	    <h1>Welcome </h1>
		</body>
	</html>
	`)
}
func main() {

	// http.HandleFunc("/login/", mylogin) // whatever the port on which our file is running we can use this route after the port
	// http.HandleFunc("/welcome", myWelcome)

	http.Handle("/login", http.HandlerFunc(mylogin))
	http.Handle("/login", http.HandlerFunc(myWelcome))
	fmt.Println("listening on port 3000.....") //convention h ki server is started
	http.ListenAndServe("localhost:3000", nil) // this will not work mtlb localhost:3000 bcoz we have not passed the handler but if you use "localhost:3000/login" to chalega

}
