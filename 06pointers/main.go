package main

import "fmt"
func main(){
	fmt.Println("pointers");

	// var age int = 22		
	//  var ptr *int;
	//  ptr = &age
	//  fmt.Println(*ptr)

	myNumber := 43;
	var ptr = &myNumber;
	fmt.Println(ptr);
	fmt.Println(*ptr);
}