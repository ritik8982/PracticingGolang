package main

import "fmt"

func main() {
	// fmt.Println("slice");
	// var fruitlist = []string{"apple","oranges"};
	// fmt.Printf("type of fruitlist is %T",fruitlist);
	// fmt.Println();
	// fruitlist = append(fruitlist, "mango","banana");
	// fmt.Println("fruitList: ",fruitlist);

	// //some operation
	// fruitlist = append(fruitlist[1:]);
	// fmt.Println("new fruitlist is: ",fruitlist);

	// fruitlist = append(fruitlist[1:3]);
	// fmt.Println("new fruitlist is: ",fruitlist);

	// fruitlist = append(fruitlist[:3]);
	// fmt.Println("new fruitlist is: ",fruitlist);

	// highScores := make([]int, 4)
	// highScores[0] = 235
	// highScores[1] = 230
	// highScores[2] = 232
	// highScores[3] = 239
	// highScores = append(highScores, 546,839);
	// fmt.Println("before sorting: ",highScores);

	// sort.Ints(highScores);
	// fmt.Println("after sorting: ",highScores);
	// fmt.Println(sort.IntsAreSorted(highScores));

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses);	

	var index int = 2;
	courses = append(courses[:index],courses[index+1:]...);
	fmt.Println(courses);

}
