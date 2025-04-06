package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 23, 15}

	fmt.Print(nums)

	maxi := nums[0]

	for _, num := range nums[1:] {
		if num > maxi {
			maxi = num
		}
	}

	fmt.Printf("\nThe maximum value in the slice is: %d\n", maxi)
}
