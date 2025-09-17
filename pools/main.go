package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Print("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second * 2)
		results <- job
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// jobs channel closed
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}

}
