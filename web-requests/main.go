package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// sample API endpoint
const url = "https://reqres.in/api/users/2"


func fetchData() {
	response, err := http.Get(url)

	if err != nil {
		println("Error fetching data:", err)
		return
	}

	fmt.Printf("response is of type: %T \n", response)
	fmt.Println(response)

	defer response.Body.Close() // caller's responsibility to close the response connection - never forget this!

	// Create a buffered reader
	reader := bufio.NewReader(response.Body)
	// fmt.Println("Response body: ", reader)


	

	for {
		// Read a chunk of data
		chunk, err := reader.ReadBytes('\n')
		if err != nil {
			// If EOF is reached, break the loop
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading response body:", err)
			return
		}
		// Print the chunk of data
		fmt.Println("this runs")
		fmt.Println(string(chunk))
	}
}

func main() {
	println("Welcome to web-requests!")
	fetchData();
}