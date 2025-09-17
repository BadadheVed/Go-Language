package main

import (
	"fmt"
	"sync"
)

func counter(id int, wg *sync.WaitGroup) {
	sum := 0
	defer wg.Done()
	for i := 1; i <= 1000000; i++ {
		sum += i
	}
	fmt.Println("Sum is ", sum)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // addin one thread
		go counter(i, &wg)
	}
	wg.Wait()
	fmt.Println("All the 3 cpu intensive tasks done")
}
