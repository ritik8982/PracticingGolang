package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello go")
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
}
