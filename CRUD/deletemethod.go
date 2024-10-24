package main

import (
	"fmt"
	"io"
	"net/http"
)

func PerformDELETERequest() {

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error creating  put request", err)
		return
	}

	//Send the request
	client := http.Client{}
	res, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Error sending put request", err1)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
	fmt.Println("Response Status: ", res.Status)
}
