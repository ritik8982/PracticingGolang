package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev";

func main() {
	fmt.Println("learning about web request")
	response,err := http.Get(url);

	if err != nil{
		panic(err);
	}
	fmt.Printf("respone is of type :%T\n",response);	//respone is of type :*http.Response

	defer response.Body.Close();		// it is the caller responsiblity to close the connection

	dataBytes,err := ioutil.ReadAll(response.Body);
	if err != nil{
		panic(err);
	}
	fmt.Println("",string(dataBytes))
}
