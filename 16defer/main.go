package main

import "fmt"
func main(){
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	fmt.Println("after second defer")
	myFunc();
	defer fmt.Println("third defer")

}

func myFunc(){
	for i:=1; i<=5; i++{
		defer fmt.Println(i);
	}
}