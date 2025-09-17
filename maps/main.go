package main

import "fmt"

func main() {
	studentGrades := make(map[string]int) // here string is the key type and the int is the value type
	studentGrades["Prince"] = 34
	studentGrades["alice"] = 35
	studentGrades["bob"] = 36
	studentGrades["charlie"] = 38
	fmt.Println(studentGrades["bob"])
	delete(studentGrades, "bob")
	fmt.Println(studentGrades["bob"])
}
