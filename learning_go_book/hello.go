package main

import "fmt"

func main() {

	fmt.Printf("Hello")

	var myLastInitial rune = 'B'
	fmt.Println(myLastInitial)

	// declaration using type inference. this can only be used inside functions
	a := 10
	println(a)

	// declaration lists. to be used outside functions
	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)

	println(x, y, z, d, e, f, g)

	// constants
	const (
		idKey   = "id"
		nameKey = "name"
	)
}
