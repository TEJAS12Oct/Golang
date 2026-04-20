package main

import "fmt"

// Main function
func main() {

	// Case 1: Use all values
	add, mul, div := operations(105, 7)
	fmt.Println("Using all values:")
	fmt.Println("105 + 7 =", add)
	fmt.Println("105 x 7 =", mul)
	fmt.Println("105 / 7 =", div)

	fmt.Println("----------------------")

	// Case 2: Ignore division
	addOnly, mulOnly, _ := operations(105, 7)
	fmt.Println("Ignoring division:")
	fmt.Println("105 + 7 =", addOnly)
	fmt.Println("105 x 7 =", mulOnly)

	fmt.Println("----------------------")

	// Case 3: Use only division
	_, _, divOnly := operations(105, 7)
	fmt.Println("Only division:")
	fmt.Println("105 / 7 =", divOnly)
}

// Function returning 3 values
func operations(n1 int, n2 int) (int, int, int) {

	add := n1 + n2
	mul := n1 * n2
	div := n1 / n2

	return add, mul, div
}
