package main

import (
	"strings"
)

func main() {
	text := `
		Hello World
		Hello World
		Hello
		`

	words := strings.Fields(text)
	counts := map[string]int{}

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	print("Word counts:\n")
	for word, count := range counts {
		println(word, count)
	}
}
