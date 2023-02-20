package main

import "fmt"

func main() {
	fmt.Println("methods in golang");

	ritik := User{"ritik","ritik@gmail.com",true,16};
	fmt.Println(ritik.Email);
	// ritik.getStatus();
	ritik.changeEmail();
	fmt.Println(ritik.Email);
}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}

//this is how you define a method of struct
func (u User) getStatus(){
	fmt.Println("IN USER ACTIVE: ",u.Status);
}

func (u User) changeEmail(){
	u.Email = "test@gmail.com";
	fmt.Println("in changeEmail = ",u.Email);
}