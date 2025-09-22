package main

import (
	"fmt"
	"sync"
)

func helper(i int, wg *sync.WaitGroup, res chan<- int, ip <-chan int) {
	defer wg.Done()
	for ips := range ip {
		//fmt.Println("Making the channel in the square", ips)
		res <- (ips * ips)
		//fmt.Println("Square done with number as", ips)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	results := make(chan int, 5)
	input := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go helper(i, &wg, results, input)
	}
	for _, num := range numbers {
		input <- num
	}
	close(input)
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println("Square:", r)
	}

	fmt.Println("✅ Done processing!")
	wg.Wait()
}
