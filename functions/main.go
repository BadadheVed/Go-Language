package main

import "fmt"

func SimpleFunc() {
	fmt.Println("Hello I am a Simple function")
}

func add(a, b int) int {
	return a + b
}

func main() {
	SimpleFunc()
	fmt.Println(add(5, 6))
}
