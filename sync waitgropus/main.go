package main

import (
	"fmt"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}

func main() {
	fmt.Println("This is the Main function")

	// intialise a variable wg of type WaitGroup

	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1) // increamented the wait grouo counter
		// this will keep the count of the GoRoutines which are executing
		go Worker(i, &wg)
	}
	wg.Wait() // wait till all the waitgropus are done executing
	fmt.Println("End of the main function")
}
