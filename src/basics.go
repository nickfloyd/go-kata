package basics

import (
	"fmt"
)

func main() {
	assignments(10, 2)
}

func assignments(x int, y int) int {
	// 1. Get the mean of two numbers
	// 2. Use the Printf to format a string with the mean
	// 3. Use the formatter %T to also print the type of the mean

	var mean int

	// Calcualte mean

	fmt.Printf("Mean: %v.  Type: %T\n", mean, mean)
	return mean
}
