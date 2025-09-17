package main

import (
	"fmt"
)

func add(a, b int) int {
	return (a + b)
}

func main() {
	defer fmt.Println(add(4, 5))
	fmt.Println("The addition is ")
}
