package basics

import (
	"fmt"
)

func main() {
	assignments(10, 2)
}

// https://golang.org/ref/spec#Variables
func assignments(x int, y int) int {
	// 1. Get the mean of two numbers
	// 2. Use the Printf to format a string with the mean
	// 3. Use the formatter %T to also print the type of the mean

	var mean int

	// Calcualte mean

	fmt.Printf("Mean: %v.  Type: %T\n", mean, mean)
	return mean
}

// https://golang.org/ref/spec#If_statements
func ifStatements(val int) string {
	// Compaire the value passed using if/else and return the following messages
	// if val is less than 5 then return "The value is SMALL"
	// if the value is between 6 and 10 the return "The value is MEDIUM"
	// if the value is greater than 10 then return "The value is LARGE"

	message := ""

	// Solve

	return message
}

// https://golang.org/ref/spec#Switch_statements
func switchStatements(val int) string {
	// Compaire the value passed using a expression switch statement without a condition and return the following messages
	// if val is less than 5 then return "The value is SMALL"
	// if the value is between 6 and 10 the return "The value is MEDIUM"
	// if the value is greater than 10 then return "The value is LARGE"

	message := ""

	// Solve

	return message
}