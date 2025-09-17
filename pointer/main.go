package main

import "fmt"

func main() {
	num := 2
	var ptr *int
	ptr = &num

	fmt.Println("The Num has value of", num)
	fmt.Println("The Num has value of", ptr)
	var ptr1 *int
	if ptr1 == nil {
		fmt.Println("The pointer is null please assign it  a value")
	}
}
