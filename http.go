package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, error := contentType("https://www.example.com")

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		print("Content-Type:", ctype)
	}
}

func contentType(url string) (string, error) {
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("Content-Type not found")
	}

	return ctype, nil
}
