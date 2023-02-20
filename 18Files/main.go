package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("learning about files")
	content := "this needs to go in a file - ritik.shirvastava@leadsquared.com"

	file, err := os.Create("./myGmail.txt") //it will create the file and return it might return error that why we are using comma ok syntax

	checkNilError(err)

	length, err := io.WriteString(file, content) // io se likh sakte hai is file me ye content dal do, daaldega to length of the file dega nhi to err
	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close() // this is recommended
	readFile("./myGmail.txt")
}

// function to read the file
func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("text data inside the file is = \n", dataByte)         // the data is comming in byte format
	fmt.Println("text data inside the file is = \n", string(dataByte)) // the data will be in string format
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
