package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	var x int
	var y int

	// used only for assignment. ':=' also declares the variable
	x = 1
	y = 2

	fmt.Printf("x = %v, of type %T\n\n", x, x)
	fmt.Printf("y = %v, of type %T\n\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean = %v, of type %T\n\n", mean, mean)

	// initialisations inside an if statement
	a := 1
	b := 2

	if c := a + b; c < 10 {
		fmt.Printf("c = %v, of type %T\n\n", c, c)
	} else {
		fmt.Println("c is greater than or equal to 10")
	}

	//	Strings are immutable in Go

	// Example of making an HTTP GET request
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		log.Fatal(err)
	}
	// Always close the response body when done
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Body length: %d bytes\n", len(body))
}
