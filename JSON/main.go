package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("This is the JSON program")
	var p1 Person
	p1.Name = "Ved"
	p1.Age = 18
	p1.IsAdult = true
	 // fmt.Println("The Person Without JSON is ", p1)
	//  encoding is done by the json.Marshal(p1) this function
	res, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("the error is ", err)
		return
	}

	println("The JSON Object is", string(res))

	// now to decode the object into string
	// we will use the function as
	var personData Person

	err = json.Unmarshal(res, &personData)
	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("The unmarshalled object is ", personData)
}
