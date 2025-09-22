package main

import (
	"fmt"
	"sync"
)

func worker(i string, res chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	res <- i
}

func tsk1() { // Two Goroutines, One Channel

	mych := make(chan string, 10)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker("Hello From GoRoutine 1", mych, &wg)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker("Hello From GoRoutine 2", mych, &wg)
	}

	go func() {
		wg.Wait()
		close(mych)
	}()

	for msg := range mych {
		fmt.Println(msg)
	}

}

func main() {
	mych := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { // go rotine that produces numebr form 1--> 10
		for i := 1; i < 11; i++ {
			mych <- i
		}
		wg.Done()
		close(mych)
	}()

	go func(res <-chan int, wg *sync.WaitGroup) { // function to calculate the square
		defer wg.Done()
		for n := range mych {
			fmt.Println("Square Of", n, "is ", n*n)
		}
	}(mych, &wg)
	wg.Wait()
	fmt.Println("✅ Done")

}
