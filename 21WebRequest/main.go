package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("learning how to make a post request")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormData();

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content) //.write will return the length(byte) mtlb same to "response.ContentLength"

	//same
	fmt.Println("byte count is : ", byteCount)
	fmt.Println("Content Length: ", response.ContentLength)

	//same
	fmt.Println("actual content : ", responseString.String()) // this will print the data like "fmt.Println(string(content))"
	fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	//1. url
	const myurl = "http://localhost:3000/post"

	// 2. json data , payload
	requestBody := strings.NewReader(`
	 {
		"courseName" :"learn golang",
		"price" : 0,
		"platform":"learnCodeOnline.in"
	 }
	 `)

	//3. post data
	response, err := http.Post(myurl, "application/json", requestBody) //thunder -> response -> header me application/json

	//4. handle error
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 5. read the data
	contentInByte, _ := ioutil.ReadAll(response.Body)

	//6. convert the data in string
	fmt.Println(string(contentInByte))
}

func PerformPostFormData() {
	//1. url
	const myurl = "http://localhost:3000/postform"

	//2. form data
	data := url.Values{}

	data.Add("firstName", "ritik")
	data.Add("firstName", "shrivastava")
	data.Add("firstName", "ritik@gmail.com")

	//3. post form data

	response, err := http.PostForm(myurl, data)

	//4. handle error
	if err != nil {
		panic(err)
	}

	//close the request
	defer response.Body.Close()

	// 5. read the data
	content, _ := ioutil.ReadAll(response.Body)

	//6. convert it into string
	fmt.Println(string(content))

}
