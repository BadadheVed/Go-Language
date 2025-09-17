package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data int = 32
	fmt.Println(data)
	var datafloat float64 = float64(data)
	datafloat += 0.2
	fmt.Println(datafloat)
	num := 123
	str := strconv.Itoa(num)
	fmt.Println(str)
}
