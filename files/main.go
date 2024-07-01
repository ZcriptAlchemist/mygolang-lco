package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to files!")

	// Reading a file

	// Step 1: Open the file
	f1, err:= os.Open("./read-file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Step 2: Create a reader
	r := bufio.NewReader(f1)

	// Step 3: Read line by line
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Println(line)
	}
	
	// Step 4: Close the file
	defer f1.Close()

	// Writing to a file

	// Step 1: Create/Open output-file
	f2, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Step 2: Convert data to []byte
	data := []byte("this text will be written into the file!\n")

	// Step 3: Write data to file
	_, err = f2.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		f2.Close()
		return
	}

	fmt.Println("Data written successfully!")

	// Step 4: Close the file
	defer f2.Close()
}