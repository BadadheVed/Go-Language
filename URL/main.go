package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("This is the URL File")
	urlstr := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("The type of the url is %T\n", urlstr)
	newURL, _ := url.Parse(urlstr)
	fmt.Printf("the type og the new url is %T\n ", newURL)
}
