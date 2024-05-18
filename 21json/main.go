package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // means this field will not be reflected to consumer of API
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json topic")

	// encoding the json
	// means you have the data (slice, or key vals or array etc)
	// & to convert the data it into a valid Json
	// EncodeJson()

	DecodeJson()

}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS bootcamp", 299, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"MERN bootcamp", 199, "Udemy", "asd123", []string{"web", "js"}},
		{"AngularJS bootcamp", 399, "Udemy", "mvc123", nil},
	}

	// package data as JSON data
	// kind of an implementation of JSON
	// pass interface (alternative version of the struct)
	// finalJson, err := json.Marshal(myCourses)

	//param1: interface or struct, p2: prefix, p3: based on what you want to indent the value
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

// consume json
func DecodeJson() {
	fmt.Println("meow")

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS bootcamp",
		"Price": 299,
		"website": "Udemy",
		"tags": [
			"web-dev",
			"js"
		]
	}
	`)

	// verify if correct JSON format
	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// An interface type is defined as a set of method signatures.
	// type Abser interface {
	// 	Abs() float64
	// }

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v & type is %T\n", k, v, v)
	}

}
