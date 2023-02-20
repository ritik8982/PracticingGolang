package main

import "fmt"

func main() {
	fmt.Println("jalwa")
	q()

	sum := adder(3, 4)
	fmt.Println("sum= ", sum)

	proSum := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("proSum = ", proSum)

	name, age := nameAndAge()
	fmt.Println("name: ", name, "age: ", age)
}

// values jo hai wo ab slice ban gaya hai || function kitne bhi argument accept kare
func proAdder(values ...int) int {

	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// return multiple thing
func nameAndAge() (string, int) {
	return "ritik", 23
}

func adder(a int, b int) int {
	return (a + b)
}

func q() {
	fmt.Println("inside the anonymous function")
}
