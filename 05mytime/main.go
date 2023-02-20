package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main")
	
	presentTime := time.Now();
	fmt.Println(presentTime);
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"));

	createDate := time.Date(2020,time.February,15,22,15,0,0,time.UTC);
	fmt.Println(createDate);
}