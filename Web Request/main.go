package main

import (
	"fmt"
	"io"
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

	data ,err := io.ReadAll(res.Body) // this function is used to get the data of the http.response;
	if err != nil {
		fmt.Println("Error is ", err)
		return
	}

	fmt.Println("Response is ", string(data))
}
