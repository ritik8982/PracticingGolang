package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex

// line no.7 ka mtlb ek thread bano jo ki greeter ko hello ke sath print (chalayega)karega, but you never said when to comeback and you never waited to comeback
func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websiteList {

		go getStatusCode(web)
		wg.Add(1) //kitne goroutine(thread) bahar hai
	}

	wg.Wait() //usually goes to the end of the method(main), because this is responsible ki main method end na ho
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // ye program ke last me hi chalega
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error in end point")
	}
	// jaha bhi read karne wale ho usse phele lock() and then unlock() kardena read karne ke baad
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
