package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	// read response

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content received is: ", content)

}
