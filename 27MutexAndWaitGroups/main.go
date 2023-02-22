package main

import (
	"fmt"
	"sync"
)

func main() {

	var score = []int{0}

	wg := &sync.WaitGroup{} //WaitGroup are write as pointers
	mut := &sync.Mutex{}

	// teen wait group hai ek ek karke add kar sakte ho mtlb wg.add(1) ya fir wq.add(3) kardo kiuki teen hi wait group hai
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")

		mut.Lock() //locking and unlocking using mutex
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut) //upar function bana aur yahi call hogaya aur ye parameter pass ho gaye usme
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("three R")

		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	wg.Wait() // we have to notify the main method that we have enable the wait group
	fmt.Println(score)
}
