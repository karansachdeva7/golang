package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	// developer's responsibility to close the request
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// One way
	// read response using ioutil
	// content is in the byte format
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(content)
	// fmt.Println(string(content))

	// Another way
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	// whenever you have data in the form of byte, you can use below syntax
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is:", byteCount)

	// response is inside responseString (holds raw data)
	// convert data inside to string format
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const url = "http://localhost:8000/post"

	// fake json payload
	// create new reading format i.e. create any kinda data
	requestBody := strings.NewReader(`
		{
			"coursename":"golang course",
			"price": 0,
			"platform":"golang dev"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	// so that request of the response is closed
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

// 
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// create form data
	data := url.Values{}
	// you can replace with variable (i.e. source of data)
	data.Add("firstname", "karan")
	data.Add("lastname", "sachdeva")
	data.Add("email", "karan@gmail.com")

	// issues post request which is url encoded
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
