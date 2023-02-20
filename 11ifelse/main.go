package main

import "fmt"

func main(){
	fmt.Println("if else");
	loginCount := 23;
	var userType string;

	if loginCount < 10{
		userType = "Normal user";
	}else if loginCount > 10 && loginCount <25{
		userType = "premium user";
	}else{
		userType = "ultra premium uesr";
	}
	fmt.Println(userType);

	if num := 3; num < 10{
		fmt.Println("the num is less than 10");
	}else{
		fmt.Println("the num is greater than or equal to 10");
	}
}