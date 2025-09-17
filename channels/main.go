package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	n := len(ch)
	for i := 0; i < n; i++ {
		x := <-ch
		fmt.Println(x)
	}

}
