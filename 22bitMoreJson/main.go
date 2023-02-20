package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string `json:"-"`
	Password string
	Tags     []string //slice of string
}

func main() {
	fmt.Println("create json data")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	//slice of courses - slice ka array hai hota hai dynamic wala, to slice of course mtlb har ek index pe ek course hoga aur apn banate samay hi initialse kar rahe hai to {} lagega na meri jaan
	lcoCourses := []course{
		{"reactjs bootcamp", 299, "learnCodeOnline.in", "abc123", []string{"webdev", "js"}},
		{"MERN bootcamp", 199, "learnCodeOnline.in", "bdd123", []string{"full-stack", "js"}},
		{"Angular bootcamp", 299, "learnCodeOnline.in", "ritik@12", nil},
	}

	//2. package (convert) this data as a json data

	// finalJson,err := json.Marshal(lcoCourses);

	/* o/p = [{"Name":"reactjs bootcamp","Price":299,"Platform":"learnCodeOnline.in","Password":"abc123","Tags":["webdev","js"]},
	{"Name":"MERN bootcamp","Price":199,"Platform":"learnCodeOnline.in","Password":"bdd123","Tags":["full-stack","js"]},
	{"Name":"Angular bootcamp","Price":299,"Platform":"learnCodeOnline.in","Password":"ritik@12","Tags":null}]*/

	// above data sahi se indenet nhi hai isliye padh nhi pa rahe hai agar chahte ho ki sahi se indent ho to use karo MarshalIndent

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // prefix, space \t mtlb tab space jaha hona chahiye space waha hoga \t ka space
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName":"Reactjs BootCamp",
		"price":299,
		"website":"LearnCodeOnline.in",
		"tags":["web-dev","js"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) // it will return true false

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
		//o/p => JSON was valid
		// main.course{Name:"Reactjs BootCamp", Price:299, Platform:"", Password:"", Tags:[]string{"web-dev", "js"}}
		// courseName was an alias for Name
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	// map[string]interface {}{"courseName":"Reactjs BootCamp", "price":299, "tags":[]interface {}{"web-dev", "js"}, "website":"LearnCodeOnline.in"}

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
	/*bus ache se print kiya hai upar wale ko
		key is courseName and value is Reactjs BootCamp and type is: string
	key is price and value is 299 and type is: float64
	key is website and value is LearnCodeOnline.in and type is: string
	key is tags and value is [web-dev js] and type is: []interface {}
	*/

}
