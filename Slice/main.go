package main

import "fmt"

func main() {
	number := []int{1, 2, 3}
	fmt.Println("numbers are", number)
	number = append(number, 1, 2, 3, 4, 5, 6)
	fmt.Println("The array after appending is :-", number)
	
}
