package main

import "fmt"

const LoginToken string = "dfasjdf";

func main() {
	fmt.Println("variable----------------------------------------------");

	var userName string = "ritik";
	fmt.Println(userName);
	fmt.Printf("variable is of type: %T \n",userName);

	var isLoggedin bool = false;
	fmt.Println(isLoggedin);
	fmt.Printf("variable is of type: %T \n",isLoggedin);
	
	// var smallFloat float32 = 255.345234234;

	var anotherVariable string;
	fmt.Println(anotherVariable);
	fmt.Printf("variable is of type %T \n",anotherVariable);

	numberOfUser := 30000;
	fmt.Println(numberOfUser);
}
