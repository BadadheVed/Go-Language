package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("This is Hello Function")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Say Hello Completed")
}

func sayGM() {
	fmt.Println("Good Morning")
}

func main() {
	fmt.Println("This is the Go Routines")
	go sayHello()
	go sayGM()
	time.Sleep(3000 * time.Millisecond)
}
