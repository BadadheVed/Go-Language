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
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}
	wg.Wait()
}
