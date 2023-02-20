package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hghba23"

func main() {
	fmt.Println("learning about urls")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme);		//https
	// fmt.Println(result.Host);
	// fmt.Println(result.Path);
	// fmt.Println(result.Port());
	// fmt.Println(result.RawQuery);

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams) // qparams jo hai wo url.values == mtlb key value pair
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	/* result of above
	Param is:  [reactjs]
	Param is:  [hghba23]
	*/

	//maanlo apke pass key value hai usse url banna chahte ho
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=ritik",
	}

	anotherUrl := partsofUrl.String() // this will convert the url into string mtlb
	fmt.Println(anotherUrl)

}
