package main

import "fmt"

func main() {
	fmt.Println("learning loop")
	days := []string{"sunday", "tueday", "wednesday", "Thursday", "saturday"}
	fmt.Println(days)

	// for loop

	// 1.
	for d:=0; d<len(days); d++{
		fmt.Println(days[d]);
	}

	// 2. range  will give you the index
	for i := range days {
		fmt.Println(days[i])
	}

	// 3. 
	for index, value := range days {
		fmt.Printf("index is %v and value is %v\n", index, value)
	}

	//4. while loop hota nhi hai but for loop me agar initialization na karo or direct checking condition laga do to for while jesa lagta hai

	n := 1;
	for n < 10{
		if n == 5{
			n++;
			continue;
		}
		if n == 8{
			goto lco
		}
		fmt.Println(n);
		n++;
	}
	fmt.Println("goto before")
	lco:
	fmt.Println("goto hit");
	

}
