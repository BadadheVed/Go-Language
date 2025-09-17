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

func addition(n int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sum := n * (n - 1) / 2
	for i := 0; i < 1000000; i++ {

	}
	fmt.Println("The addition thread done")
	return sum
}

func sq(n int, wg *sync.WaitGroup) int {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {

	}
	fmt.Println("The Square thread done")
	return n * n
}

func div(n int, wg *sync.WaitGroup) int {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {

	}
	fmt.Println("The Division thread done")
	return n / 1
}

func main() {
	var wg sync.WaitGroup
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n) // take input from user and store in n
	fmt.Println("You entered:", n)
	wg.Add(3)

	go div(n, &wg)
	go addition(n, &wg)
	go sq(n, &wg)
	fmt.Println("All the 3 threads completed")
	wg.Wait()
}
