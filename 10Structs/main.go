package main

import "fmt"

func main() {
	fmt.Print("learning structs\n")

	//this is how you assign value to struct
	ritik := User{"ritik", "ritik.shri@go.dev", false, 21}

	fmt.Println(ritik) //{ritik ritik.shri@go.dev false 21}

	fmt.Printf("ritik details are: %+v", ritik) //ritik details are: {Name:ritik Email:ritik.shri@go.dev Status:false Age:21}

	//this is how you access a single value
	fmt.Printf("name is %v\n", ritik.Name)
}

// this is how you define struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
