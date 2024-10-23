package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learn Web Request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting response")
		return
	}
	defer res.Body.Close()
	// fmt.Printf("Type of response is %T\n",res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response")
		return
	}
	fmt.Println("Response is ", string(data))
	handleUrl()
}

/*
Web Request:
-> A Web Request is a HTTP Request made to the server.
-> In Go, we can make web requests using the net/http package, which provides functions to create and send http requests,
as well as handle responses.
->

*/
