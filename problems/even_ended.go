package main

import "fmt"

func main() {

	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ { // Avoid duplicate products

			product := i * j
			string_product := fmt.Sprintf("%d", product)
			if string_product[0] == string_product[len(string_product)-1] {
				count++
			}
		}
	}

	fmt.Printf("Count of products with even-ended palindromic numbers: %d\n", count)
}
