package main

// gin is not install facing some issue
import (
	"fmt"
	"net/http"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
	   <head>
	   Hi
	   <head>
	   <body>
	    <h1> Ritik Shrivastava </h1>
		</body>
	</html>
	`)
}
func main() {

	http.ListenAndServe("localhost:3000", http.HandlerFunc(myfunc))

}
