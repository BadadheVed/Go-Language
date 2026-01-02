package main

import (
	"fmt"
	"net/http"
	"sync"
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

type Site struct {
	URL string
}
type Result struct {
	Status int
}

func crawl(w int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for site := range jobs {
		fmt.Println("worker id is ", w)

		res, err := http.Get(site.URL)
		if err != nil {
			fmt.Printf("Got the error while getting data from the site\n")
			continue
		}

		results <- Result{Status: res.StatusCode}
	}
}

func main() {
	// jobs := make(chan int, 5)
	// results := make(chan int, 5)

	// for w := 1; w <= 3; w++ {
	// 	go worker(w, jobs, results)
	// }
	// // jobs channel closed
	// for i := 1; i <= 5; i++ {
	// 	jobs <- i
	// }
	// close(jobs)
	// for a := 1; a <= 5; a++ {
	// 	fmt.Println("Result:", <-results)
	// }

	cha1 := make(chan Site, 3)   // jobs
	cha2 := make(chan Result, 3) // results
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go crawl(w, cha1, cha2, &wg)
	}

	// Send jobs
	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com",
		"https://google.com",
	}
	for _, url := range urls {
		cha1 <- Site{URL: url}
	}
	close(cha1) // No more jobs

	// Close results after workers are done
	go func() {
		wg.Wait()
		close(cha2)
	}()

	// Collect results
	for res := range cha2 {
		fmt.Println(res)
	}
}
