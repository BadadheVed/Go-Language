package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGET() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("The error is", err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Println("The response in is invalid")
		return
	}
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("The error is", err)
	// 	return
	// }

	// fmt.Println("the data is ", string(data))

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("the error is ", err)
		return
	}

	fmt.Println("todos", todo)
}

func performPOST() {
	todo := Todo{
		UserID:    23,
		Title:     "Ved Baddhe",
		Completed: true,
	}

	// convert todo to json

	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Print("The error is ", err)
		return
	}

	jsonString := string(data)
	// convert string to readers

	jsonReader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("errors are", err)
		return
	}

	defer res.Body.Close()

	// make the response readable
	ans, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("the response is", err)
	}
	fmt.Println("JSON DATA", string(ans))
}

func performUPDATE() {
	todo := Todo{
		UserID:    23,
		Title:     "Ved Baddhe",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("the response is", err)
		return
	}

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// create put request

	jsonString := string(jsonData)
	// convert string to readers

	jsonReader := strings.NewReader(jsonString)
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error Sending request", err)
	}

	req.Header.Set("Content-type", "application/json")
	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending request", err)
	}

	defer res.Body.Close()
	ans, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("the response is", err)
	}
	fmt.Println("JSON DATA", string(ans))
	fmt.Println("The status code is ", res.StatusCode)
}

func PerformDelete() {
	
}

func main() {
	fmt.Println("Learining CRUD Operations in go lang")

	performUPDATE()
}
