package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang ")

	mych := make(chan int) //kibal ekhi value add karo is channel me
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
 
		val,isChannelOpen := <-mych
		if isChannelOpen {

			fmt.Println(val) //listening(consumer) <-mych
		}
		
		wg.Done()

	}(mych, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		// mych <- 0
		close(mych)
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
