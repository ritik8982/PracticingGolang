package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //whatever reader variable has read it we have to store it somewhere that's where comma(err) ok sentence came in
	fmt.Println("enter the rating for our pizza: ")

	input, _ := reader.ReadString('\n') // agar sahi hoga to input me answer mil jayega agar glt hoga toh error me answer aa jayega

	//input reader me hai reader ko read ReadString karega but kaha thak read kare toh \n thak karna ye bata raha hai
	fmt.Println("thank you for giving me rating", input)
	fmt.Printf("type of this string is %T", input)
}
