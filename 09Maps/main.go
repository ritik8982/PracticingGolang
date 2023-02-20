package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	// how to declare map
	languages := make(map[string]string)

	// how to assing value
	languages["RB"] = "ruby"
	languages["py"] = "python"
	languages["js"] = "javaScript"

	// how to print and how to access
	fmt.Println("list of all languages: ", languages)
	fmt.Println("js is short form for : ", languages["js"])

	// how to delete the
	delete(languages,"RB");
	fmt.Println(languages);

	//loop through the map
	for key,value := range languages{
		fmt.Printf("for key %v, value is %v \n",key,value);
	}

}
