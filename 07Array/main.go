package main

import "fmt"

func main(){
	fmt.Println("array");
	var fruitList [4] string;		//string arr[4]  &&  string fruitlist [] = new string[4]
	fruitList[0] = "apple";
	fruitList[1] = "banana";
	fruitList[3] = "oranges";

	fmt.Println("fruitlist: ",fruitList);
	fmt.Println("length of the fruitlist: ",len(fruitList));

	var vegList = [3]string{"potato","beans","mushroom"};
	fmt.Println("vegList: ",vegList);





	// for(i:=0; i<4; i++)
	// fmt.Println(frufruitList[i]);
}