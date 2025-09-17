package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for jobs := range jobs {
		fmt.Println("Received job with value", jobs)
		time.Sleep(time.Second * 2)
		results <- (jobs * (jobs - 1) / 2)
		fmt.Println("job done with number as", jobs)
	}
}
func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		jobs <- arr[i]
	}
	close(jobs)
	for i := 0; i < len(arr); i++ {
		fmt.Println("Result:", <-results)
	}
}
