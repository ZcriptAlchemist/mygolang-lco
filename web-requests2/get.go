package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "http://localhost:8000/get"

func main() {
	fmt.Println("Hello, this is get request!")
	PerformGetRequest(url)
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer response.Body.Close()
	
	// there are two ways to handle the response body
	// 1st method using ioutil.ReadAll
	fmt.Println("Status: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content,_ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))

	// 2nd method using strings.Builder
	var responseString strings.Builder
	content2, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content2)

	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Content: ", responseString.String())
}