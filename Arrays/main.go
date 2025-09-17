package main

import "fmt"

func main() {
	fmt.Println("This is the arrays chapter")
	var age [10]int
	fmt.Println("Declared an aray of the size 10")
	age[0] = 14
	age[2] = 34
	fmt.Println("The age of the 2nd person is", age[1])
	fmt.Println("Length Of the array is ", len(age))
	name := [5]string{"Hello", "i", "am", "a", "devloper"}
	fmt.Println(name)
	arr := []int{1, 2, 3, 4}
	arr = append(arr, 45);
	fmt.Println(arr)
}
