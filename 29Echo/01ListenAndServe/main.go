package main

import (
	"fmt"
	"net/http"
)

var (
	abc string
	bcd string
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type MyWebServerType bool

// bcoz responseWriter and Request is inside the http package
func (m MyWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "well,hello there!!!") // this method take first argu is responseWriter and second argu as msg, it will print on the console
	fmt.Fprintf(w,"request is: %+v",r);
}
func main() {
	var k MyWebServerType
	http.ListenAndServe("localhost:3000", k) //frist argu is port, second is handler , handler will call ServeHTTP internally

	//404 is the default response , when you're not passing any handler
	//handler is the one which handle the http request
}

