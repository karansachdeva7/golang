package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjbnsswe"

func main() {
	fmt.Println("Welcome to handling urls in golang")

	// net is not a library but a module
	//fmt.Println(myurl)

	// parsing
	// receive error if url is malformed / some error
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// stores qparams in a better format
	qparams := result.Query()
	// url.Values (i.e. key value pairs)
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// construct url from chunks/parts of info
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=akshay",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
