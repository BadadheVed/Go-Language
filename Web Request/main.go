package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning the web requests")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("the error is ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("The type of the response is %T\n", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error is ", err)
		return
	}

	fmt.Println("Response is ", string(data))
}
