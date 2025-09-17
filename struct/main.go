package main

import "fmt"

type College struct {
	name    string
	id      int
	add     string
	student Students
}
type Students struct {
	name     string
	rollno   int
	isPasses bool
}

func main() {
	var cc College
	cc.name = "PICt"
	cc.id = 4005
	cc.add = "pune"
	cc.student.name = "ved"
	cc.student.rollno = 21

	fmt.Println("the students in the colege are", cc.student)
	
}
