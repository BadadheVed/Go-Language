package main

import (
	"fmt"
	"sync"
)

//	func worker(id int, jobs <-chan int, results chan<- int) {
//		for jobs := range jobs {
//			fmt.Println("Received job with value", jobs)
//			time.Sleep(time.Second * 2)
//			results <- (jobs * (jobs - 1) / 2)
//			fmt.Println("job done with number as", jobs)
//		}
//	}
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker ", id, "is running")
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	fmt.Println("worker ", id, "is finishd")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("working of all threads done")
}
