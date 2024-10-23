package main

import (
	"fmt"
	"net/url"
)

func handleUrl() {
	fmt.Println("Lets Learn URL")
	myUrl := "https://jsonplaceholder.typicode.com/todos/1?key1=value1&key2=value2"
	fmt.Printf("Type of URL : %T\n", myUrl)

	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error Parsing Url")
		return
	}
	fmt.Printf("Type of Parsed Url %T\n", parsedURL)
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	//Modifying URL Components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "user=kapil"

	newUrl := parsedURL.String()

	fmt.Println("New URL:", newUrl)
}

/*
In Go, the net/url package provides functionalities to parse, construct, and manipulate URLs( Uniform Resource Locators).
URLs are used to identify resources on the internet, such as web pages, images, and files.
-> Parsing URLs : The url.Parse function is used to parse a string into a URL object (url.URL struct). This allows us to break down
a URL into its individual components, such as scheme, host, path, and query parameters.
-> Accessing URL Components : Once you have a URL Object, you can access its component using various fields:
 * Scheme: Indicates the Protocol used (e.g. "http", "https").
 * Host: Specifies the domain name and optionally the port number.
 * PathL Represents the path component of the URL, which specifies the resource's location on the server.
 * RawQuery: Contains the raw query string, including query parameters.

 -> Query Parameters: query paramters are key-value pairs appended to the end of a URL, usually starting with a ? and separated by
    &. The url.Values type is used to represent query parameters as a map.

-> 	We use String method to convert the modified ur.URL object back to a string representation.
*/
