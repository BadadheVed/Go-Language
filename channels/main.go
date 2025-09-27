package main

import (
	"fmt"
	"sync"
)

func main() {
	dataChan := make(chan int)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				dataChan <- i
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for v := range dataChan {
		fmt.Println(v)
	}

}
